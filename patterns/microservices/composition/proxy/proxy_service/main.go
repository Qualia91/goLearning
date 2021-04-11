package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Proxy handler will receive a get request with url and forward on traffic to service that can deal with it
// Url to services:
// - "/hello" - hello_service (9001)
// - "/world" - world_service (9002)
type ProxyHandler struct {
}

func main() {

	ph := new(ProxyHandler)

	http.ListenAndServe(":9000", ph)

}

// Serve HTTP Function to Implement RESTfull API
func (h ProxyHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:

		path := r.URL.Path
		switch path {
		case "/hello":
			callService("9001", rw)
		case "/world":
			callService("9002", rw)
		default:
			rw.WriteHeader(http.StatusNotFound)
		}

	default:
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func callService(port string, rw http.ResponseWriter) {
	newReader1 := strings.NewReader("my request 1")

	resp1, err := http.Post(fmt.Sprintf("http://localhost:%s", port), "application/text", newReader1)
	if err != nil {
		fmt.Printf("Error on post to %s: %s\n", port, err)
	}

	buf1, err := ioutil.ReadAll(resp1.Body)
	if err != nil {
		fmt.Printf("Error on read of response from %s: %s\n", port, err)
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write(buf1)
}
