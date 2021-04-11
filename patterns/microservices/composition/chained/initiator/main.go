package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// This initiator will call service1, which will call service 2, which will call service 3, then responses will
// be chained back up to here and presented
type InitHandler struct {
}

func main() {

	ih := new(InitHandler)

	http.ListenAndServe(":9000", ih)

}

// Serve HTTP Function to Implement RESTfull API
func (h InitHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		newReader1 := strings.NewReader("my request 1")

		resp1, err := http.Post("http://localhost:9001", "application/text", newReader1)
		if err != nil {
			fmt.Printf("Error on post to 9001: %s\n", err)
		}

		buf1, err := ioutil.ReadAll(resp1.Body)
		if err != nil {
			fmt.Printf("Error on read of response from 9001: %s\n", err)
		}

		rw.WriteHeader(http.StatusOK)
		rw.Write(buf1)

	default:
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}
