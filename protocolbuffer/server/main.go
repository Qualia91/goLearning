package main

import (
	"fmt"
	"net/http"
	pb "protocolbuffer/proto"

	"google.golang.org/protobuf/proto"
)

type Handler struct {
}

func main() {

	h := new(Handler)
	http.ListenAndServe("localhost:9001", h)

}

// Serve HTTP Function to Implement RESTfull API
func (h Handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		addrBook := &pb.AddressBook{}
		p := pb.Person{
			Id:    1234,
			Name:  "John Doe",
			Email: "jdoe@example.com",
			Phones: []*pb.Person_PhoneNumber{
				{Number: "555-4321", Type: pb.Person_HOME},
			},
		}
		addrBook.People = append(addrBook.People, &p)
		out, err := proto.Marshal(addrBook)
		if err != nil {
			fmt.Printf("error while marshalling address book: %s", err)
			return
		}

		rw.WriteHeader(http.StatusOK)
		rw.Write(out)
	default:
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}
