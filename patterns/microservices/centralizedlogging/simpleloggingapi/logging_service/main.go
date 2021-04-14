package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Handler struct {
}

func main() {

	h := new(Handler)

	err := http.ListenAndServe(":9000", h)
	if err != nil {
		panic(err)
	}

}

// Serve HTTP Function to Implement RESTfull API
func (h Handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:

		buf, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("Error on read of request from 9000: %s\n", err)
		}

		// would save to database here, but im going to just print it out to show it working
		fmt.Println(string(buf))

	default:
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}
