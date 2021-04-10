package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

type webpageHandler struct{}

func main() {
	wh := new(webpageHandler)

	http.Handle("/", wh)

	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatal(err)
	}
}

func (wh webpageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		resp, err := http.Post("http://localhost:3001", "application/text", bytes.NewBuffer([]byte("Hello World")))
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		w.Write(body)
	}

}
