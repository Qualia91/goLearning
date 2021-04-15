package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	pb "protocolbuffer/proto"

	"google.golang.org/protobuf/proto"
)

type Handler struct {
}

func main() {

	h := new(Handler)
	http.ListenAndServe("localhost:9000", h)

}

// Serve HTTP Function to Implement RESTfull API
func (h Handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:

		byteArray := []byte("hello")
		buf := bytes.NewBuffer(byteArray)
		respBuf, err := http.Post("http://localhost:9001", "Request", buf)
		if err != nil {
			fmt.Printf("error with post response: %s", err)
			return
		}

		resp, err := ioutil.ReadAll(respBuf.Body)
		if err != nil {
			fmt.Printf("error in reading response buffer: %s", err)
			return
		}

		// unmarshal protobuff
		book := &pb.AddressBook{}
		if err = proto.Unmarshal(resp, book); err != nil {
			fmt.Printf("error unmarshalling protobuff: %s", err)
			return
		}

		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(book.People[0].Name))

	default:
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}
