package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"microservices/centralizedlogging/simpleloggingapi/log_lib"
	"net/http"
	"os"
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

		// create log using shared library log_lib
		lm := log_lib.LogData{log_lib.DEBUG, 1, 1, h.port, "Log message"}

		data, err := json.Marshal(lm)
		if err != nil {
			fmt.Printf("Error converting log message to json: %s\n", err)
		}

		buffer := bytes.NewBuffer(data)

		_, err = http.Post("http://localhost:9000", "application/json", buffer)
		if err != nil {
			fmt.Printf("Error on post to logging service: %s\n", err)
		}

	default:
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}
