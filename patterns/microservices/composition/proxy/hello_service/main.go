package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
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

		messageFromInitiator := string(buf) + " hello service"

		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(messageFromInitiator))

	default:
		println("default")
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}

}
