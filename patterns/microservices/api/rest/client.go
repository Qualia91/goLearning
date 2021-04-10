package main

import "net/http"

type ClientHandler struct{}

func main() {

	ch := new(ClientHandler)

	http.ListenAndServe(":9000", ch)

}

func (cl ClientHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodConnect:
		println("http.MethodConnect")
		rw.WriteHeader(http.StatusOK)
	case http.MethodGet:
		println("http.MethodGet")
		rw.WriteHeader(http.StatusOK)
	case http.MethodPost:
		println("http.MethodPost")
		rw.WriteHeader(http.StatusOK)
	case http.MethodPut:
		println("http.MethodPut")
		rw.WriteHeader(http.StatusOK)
	case http.MethodPatch:
		println("http.MethodPatch")
		rw.WriteHeader(http.StatusOK)
	case http.MethodDelete:
		println("http.MethodDelete")
		rw.WriteHeader(http.StatusOK)
	case http.MethodHead:
		println("http.MethodHead")
		rw.WriteHeader(http.StatusOK)
	case http.MethodOptions:
		println("http.MethodOptions")
		rw.WriteHeader(http.StatusOK)
	case http.MethodTrace:
		println("http.MethodTrace")
		rw.WriteHeader(http.StatusOK)
	default:
		println("default")
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}

}
