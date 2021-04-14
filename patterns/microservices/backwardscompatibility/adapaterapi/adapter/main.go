package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Handler struct{}

func main() {
	ch := new(Handler)
	http.ListenAndServe(":9001", ch)
}

func (cl Handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodPost:

		buf, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("Error on read of request from 9000: %s\n", err)
		}

		messageFromInitiator := string(buf)

		// this symbolizes the client side data structure being transformed to fit the backend data structure
		newReader1 := strings.NewReader(messageFromInitiator + " converted to backend struct")

		resp1, err := http.Post("http://localhost:9002", "application/text", newReader1)
		if err != nil {
			fmt.Printf("Error on post to 9002: %s\n", err)
		}

		buf, err = ioutil.ReadAll(resp1.Body)
		if err != nil {
			fmt.Printf("Error on read of response from 9002: %s\n", err)
		}

		// this symbolizes the backend data being transfromed back into the client data struct
		returnMessage := string(buf) + " and now back to client struct"

		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(returnMessage))

	default:
		println("default")
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}

}
