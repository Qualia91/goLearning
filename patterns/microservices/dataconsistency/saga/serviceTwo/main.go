package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// this one will fail commit
type ClientHandler struct{}

func main() {
	ch := new(ClientHandler)
	http.ListenAndServe(":9002", ch)
}

func (cl ClientHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodPost:

		buf, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("Error on read of request from 9000: %s\n", err)
		}

		messageFromInitiator := string(buf)

		switch messageFromInitiator {
		case "COMMIT":
			fmt.Println("Trying to commit but failed")
			rw.WriteHeader(http.StatusOK)
			rw.Write([]byte("N"))
		case "ROLLBACK":
			fmt.Println("Rollback Data")
		}

	default:
		println("default")
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}

}
