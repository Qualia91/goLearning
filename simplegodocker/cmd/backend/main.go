package main

import (
	"log"
	"net/http"
)

func main() {
	// register handler used in app
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world from backend"))
	})

	err := http.ListenAndServe(":3002", nil)

	if err != nil {
		log.Fatal(err)
	}
}
