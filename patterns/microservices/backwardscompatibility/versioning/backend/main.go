package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Handler struct{}

func main() {
	ch := new(Handler)
	http.ListenAndServe(":9001", ch)
}

func (cl Handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodPost:

		versionStr := r.Header.Get("Version")
		fmt.Println(versionStr)

		buf, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("Error on read of request from 9000: %s\n", err)
		}

		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(string(buf) + " was of version " + versionStr))

	default:
		println("default")
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}

}
