package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// TransactionManager handler will receive a get request to start the transaction.
// Url to services:
// - "/hello" - hello_service (9001)
// - "/world" - world_service (9002)
type TMHandler struct {
}

func main() {

	h := new(TMHandler)

	http.ListenAndServe(":9000", h)

}

// Serve HTTP Function to Implement RESTfull API
func (h TMHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:

		respOne := prepareService("9001")
		respTwo := prepareService("9002")

		if respOne && respTwo {
			// both service can save data, to issue a commit message
			sendCommit("9001", rw)
			sendCommit("9002", rw)
		} else {
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

func sendCommit(port string, rw http.ResponseWriter) {
	newReader1 := strings.NewReader("COMMIT")
	_, err := http.Post(fmt.Sprintf("http://localhost:%s", port), "application/text", newReader1)
	if err != nil {
		fmt.Printf("Error on post to %s: %s\n", port, err)
	}
}

func prepareService(port string) bool {
	newReader1 := strings.NewReader("PREPARE")

	resp1, err := http.Post(fmt.Sprintf("http://localhost:%s", port), "application/text", newReader1)
	if err != nil {
		fmt.Printf("Error on post to %s: %s\n", port, err)
		return false
	}

	buf, err := ioutil.ReadAll(resp1.Body)
	if err != nil {
		fmt.Printf("Error on read of response from %s: %s\n", port, err)
		return false
	}

	respString := string(buf)

	return respString == "Y"

}
