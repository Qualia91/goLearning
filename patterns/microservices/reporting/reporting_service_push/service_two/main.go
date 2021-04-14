package main

import (
	"fmt"
	"net/http"
	"strings"
)

type ClientHandler struct{}

func main() {
	ch := new(ClientHandler)
	http.ListenAndServe(":9002", ch)
}

func (cl ClientHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:

		newReader1 := strings.NewReader("9002:Request Report")

		_, err := http.Post("http://localhost:9000", "application/text", newReader1)
		if err != nil {
			fmt.Printf("Error on post to 9000: %s\n", err)
			return
		}

	default:
		println("default")
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}

}
