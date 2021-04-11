package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// This handler will receive get from user, send posts to
// multiple services, and aggregate the responses from the services
// and present back to the user
type AggHandler struct {
}

func main() {

	ah := new(AggHandler)

	http.ListenAndServe(":9000", ah)

}

// Serve HTTP Function to Implement RESTfull API
func (h AggHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		newReader1 := strings.NewReader("my request 1")
		newReader2 := strings.NewReader("my request 2")
		newReader3 := strings.NewReader("my request 3")
		resp1, err := http.Post("http://localhost:9001", "application/text", newReader1)
		if err != nil {
			fmt.Printf("Error on post to 9001: %s\n", err)
		}

		resp2, err := http.Post("http://localhost:9002", "application/text", newReader2)
		if err != nil {
			fmt.Printf("Error on post to 9002: %s\n", err)
		}

		resp3, err := http.Post("http://localhost:9003", "application/text", newReader3)
		if err != nil {
			fmt.Printf("Error on post to 9003: %s\n", err)
		}

		buf1, err := ioutil.ReadAll(resp1.Body)
		if err != nil {
			fmt.Printf("Error on read of response from 9001: %s\n", err)
		}

		buf2, err := ioutil.ReadAll(resp2.Body)
		if err != nil {
			fmt.Printf("Error on read of response from 9002: %s\n", err)
		}

		buf3, err := ioutil.ReadAll(resp3.Body)
		if err != nil {
			fmt.Printf("Error on read of response from 9003: %s\n", err)
		}

		rw.WriteHeader(http.StatusOK)
		rw.Write(buf1)
		rw.Write(buf2)
		rw.Write(buf3)
	default:
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}
