package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// This client will call the adapter service, which will transform the data sent to the internal model, which sends it
// to the backend. The ensures backwards compatibility as this client data struct can be changed, and all that will need to
// change is the conversion in the api. And vice verse
type Handler struct {
}

func main() {

	h := new(Handler)

	http.ListenAndServe(":9000", h)

}

// Serve HTTP Function to Implement RESTfull API
func (h Handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		newReader1 := strings.NewReader("my request 1")

		// create client
		client := http.DefaultClient

		req, err := http.NewRequest("POST", "http://localhost:9001", newReader1)
		if err != nil {
			fmt.Printf("Error on creating request: %s", err)
		}
		req.Header.Set("Version", "1.0.1")

		resp1, err := client.Do(req)
		if err != nil {
			fmt.Printf("Error on post to 9001: %s\n", err)
		}

		buf1, err := ioutil.ReadAll(resp1.Body)
		if err != nil {
			fmt.Printf("Error on read of response from 9001: %s\n", err)
		}

		rw.WriteHeader(http.StatusOK)
		rw.Write(buf1)

	default:
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}
