package main

import (
	"fmt"
	"net/http"
	"strings"
)

// Client service receives get from user and passes to API gateway which access the API backend for user
type ClientHandler struct {
}

func main() {

	ph := new(ClientHandler)

	http.ListenAndServe(":9000", ph)

}

// Serve HTTP Function to Implement RESTfull API
func (h ClientHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:

		newReader1 := strings.NewReader("Client request")

		_, err := http.Post("http://localhost:9001", "application/text", newReader1)
		if err != nil {
			fmt.Printf("Error on post to 9001: %s\n", err)
		}

	default:
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}
