package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type ClientHandler struct{}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ch := new(ClientHandler)
		http.ListenAndServe(fmt.Sprintf(":%s", scanner.Text()), ch)
	}

}

func (cl ClientHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodPost:

		buf, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("Error on read of request from 9000: %s\n", err)
		}

		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(fmt.Sprintf("Response to %s", buf)))

	default:
		println("default")
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}

}
