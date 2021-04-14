package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Handler struct{}

func main() {
	ch := new(Handler)
	http.ListenAndServe(":9002", ch)
}

func (cl Handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodPost:

		buf, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("Error on read of request from 9001: %s\n", err)
		}

		rw.WriteHeader(http.StatusOK)
		rw.Write(buf)

	default:
		println("default")
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}

}
