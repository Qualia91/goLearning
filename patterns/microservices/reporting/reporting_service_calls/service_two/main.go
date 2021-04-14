package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type ClientHandler struct{}

func main() {
	ch := new(ClientHandler)
	http.ListenAndServe(":9002", ch)
}

func (cl ClientHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodPost:

		_, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("Error on read of request from 9001: %s\n", err)
		}

		time.Sleep(5 * time.Second)

		rw.Write([]byte("Report from client reporting service 9002"))

	default:
		println("default")
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}

}
