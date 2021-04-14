package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

// Service that sends log messages to the log service
type Handler struct {
	port string
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		h := &Handler{port: ":" + scanner.Text()}
		http.ListenAndServe(fmt.Sprintf(":%s", scanner.Text()), h)
	}

}

// Serve HTTP Function to Implement RESTfull API
func (h Handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:

		newReader1 := strings.NewReader(fmt.Sprintf("Client log message from %s", h.port))

		_, err := http.Post("http://localhost:9000", "application/text", newReader1)
		if err != nil {
			fmt.Printf("Error on post to logging service: %s\n", err)
		}

	default:
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}
