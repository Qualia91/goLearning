package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type ClientHandler struct{}

func main() {
	ch := new(ClientHandler)
	http.ListenAndServe(":9001", ch)
}

func (cl ClientHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodPost:

		buf, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("Error on read of request from 9000: %s\n", err)
		}

		messageFromInitiator := string(buf)

		newReader1 := strings.NewReader(messageFromInitiator + " append service 1 data")

		resp1, err := http.Post("http://localhost:9002", "application/text", newReader1)
		if err != nil {
			fmt.Printf("Error on post to 9002: %s\n", err)
		}

		buf1, err := ioutil.ReadAll(resp1.Body)
		if err != nil {
			fmt.Printf("Error on read of response from 9002: %s\n", err)
		}

		rw.WriteHeader(http.StatusOK)
		rw.Write(buf1)

	default:
		println("default")
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}

}
