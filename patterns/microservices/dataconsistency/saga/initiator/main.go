package main

import (
	"fmt"
	"net/http"
	"strings"
)

type InitHandler struct {
}

func main() {

	h := new(InitHandler)

	http.ListenAndServe(":9000", h)

}

// Serve HTTP Function to Implement RESTfull API
func (h InitHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:

		newReader1 := strings.NewReader("Saga Request")
		_, err := http.Post("http://localhost:9009", "application/text", newReader1)
		if err != nil {
			fmt.Printf("Error on post to 9009: %s\n", err)
		}

	default:
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}
