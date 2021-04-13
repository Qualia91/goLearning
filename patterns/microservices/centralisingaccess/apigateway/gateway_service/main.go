package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// API gateway handler (9001) will receive a post request from the user service, authenticate it and pass it onto backend api services:
// - "/hello" - hello_service (9002)
// - "/world" - world_service (9003)
type GatewayHandler struct {
}

func main() {

	ph := new(GatewayHandler)

	err := http.ListenAndServe(":9001", ph)
	if err != nil {
		panic(err)
	}

}

// Serve HTTP Function to Implement RESTfull API
func (h GatewayHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:

		buf, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("Error on read of request from 9001: %s\n", err)
		}
		// add text onto body to "authenticate it"
		authenticatedString := "Authenticated string : " + string(buf)

		callService("9002", rw, authenticatedString)
		callService("9003", rw, authenticatedString)

	default:
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func callService(port string, rw http.ResponseWriter, as string) {
	newReader1 := strings.NewReader(as)

	_, err := http.Post(fmt.Sprintf("http://localhost:%s", port), "application/text", newReader1)
	if err != nil {
		fmt.Printf("Error on post to %s: %s\n", port, err)
	}

}
