package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// This handler speaks to 2 services. One will return straight away, one will hang forever.
// This handler will employ a timeout to cancel the read of the hanging service to save the application
type Handler struct {
}

func main() {

	ah := new(Handler)

	http.ListenAndServe(":9000", ah)

}

// Serve HTTP Function to Implement RESTfull API
func (h Handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		callServiceWithTimeout("ResquestOne", "9001", rw)
		callServiceWithTimeout("ResquestOne", "9000", rw)
	default:
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func callServiceWithTimeout(requestText, port string, rw http.ResponseWriter) {
	// create channel to manage timeout
	channel := make(chan []byte, 1)

	// start send and receive in go func and pass return string to channel
	go func(channel chan []byte) {
		newReader := strings.NewReader(requestText)
		resp, err := http.Post(fmt.Sprintf("http://localhost:%s", port), "application/text", newReader)
		if err != nil {
			fmt.Printf("Error on post to %s: %s\n", port, err)
		}

		buf, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error on read of response from %s: %s\n", port, err)
		}

		channel <- buf
	}(channel)

	// use select statement on channel to impl actual timeout
	select {
	case returnString := <-channel:
		rw.WriteHeader(http.StatusOK)
		rw.Write(returnString)
	case <-time.After(1 * time.Second):
		fmt.Printf("Service on port %s did not return in time\n", port)
	}

}
