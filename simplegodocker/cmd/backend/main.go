package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

type backendHandler struct{}

func main() {
	// register handler used in app

	bh := new(backendHandler)

	http.Handle("/", bh)

	err := http.ListenAndServe(":3001", nil)

	if err != nil {
		log.Fatal(err)
	}
}

func (bh backendHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodPost:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		returnStr := "Echo from server: " + string(body)

		w.Write([]byte(returnStr))
	}

}
