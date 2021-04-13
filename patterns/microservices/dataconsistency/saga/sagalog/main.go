package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// log of messages
type SagaLog struct {
	log []string
}

var sagaLog SagaLog = SagaLog{
	log: make([]string, 0),
}

// wraps up functionality of saga log and sec
// Massively oversimplified for ease of explanation
type SagaHandler struct {
}

func main() {

	h := new(SagaHandler)

	http.ListenAndServe(":9009", h)

}

// Serve HTTP Function to Implement RESTfull API
func (h SagaHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:

		// add request to saga log
		sagaLog.log = append(sagaLog.log, "New Get Request")
		fmt.Println("New Get Request")

		successOne := sendCommit("9001", rw)
		successTwo := sendCommit("9002", rw)

		// add responses to saga log
		sagaLog.log = append(sagaLog.log, fmt.Sprintf("9001 commit response: %t\n", successOne))
		fmt.Printf("9001 commit response: %t\n", successOne)
		sagaLog.log = append(sagaLog.log, fmt.Sprintf("9002 commit response: %t\n", successTwo))
		fmt.Printf("9002 commit response: %t\n", successTwo)

		if !(successOne && successTwo) {
			// add rollback to saga log
			sagaLog.log = append(sagaLog.log, "Rolling back")
			fmt.Println("Rolling back")

			// one or both cant so issue a rollback command
			sendRollback("9001", rw)
			sendRollback("9002", rw)
		}

	default:
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func sendRollback(port string, rw http.ResponseWriter) {
	newReader1 := strings.NewReader("ROLLBACK")
	_, err := http.Post(fmt.Sprintf("http://localhost:%s", port), "application/text", newReader1)
	if err != nil {
		fmt.Printf("Error on post to %s: %s\n", port, err)
	}
}

func sendCommit(port string, rw http.ResponseWriter) bool {
	newReader1 := strings.NewReader("COMMIT")
	resp, err := http.Post(fmt.Sprintf("http://localhost:%s", port), "application/text", newReader1)
	if err != nil {
		fmt.Printf("Error on post to %s: %s\n", port, err)
	}

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error on read of response from %s: %s\n", port, err)
		return false
	}

	respString := string(buf)

	return respString == "Y"

}
