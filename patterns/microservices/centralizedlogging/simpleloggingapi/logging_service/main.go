package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

type Handler struct {
	mutex *sync.RWMutex
	logs  []byte
}

func main() {

	h := &Handler{
		mutex: &sync.RWMutex{},
		logs:  make([]byte, 0),
	}

	err := http.ListenAndServe(":9000", h)
	if err != nil {
		panic(err)
	}

}

// Serve HTTP Function to Implement RESTfull API
func (h *Handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:

		buf, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("Error on read of request from 9000: %s\n", err)
		}

		// would save to database here, but im going to just print it out to show it working
		fmt.Println(string(buf))
		h.mutex.Lock()
		h.logs = append(h.logs, buf...)
		h.mutex.Unlock()

	case http.MethodGet:

		rw.WriteHeader(http.StatusOK)
		rw.Header().Set("Content-Type", "application/json")
		h.mutex.RLock()
		rw.Write(h.logs)
		h.mutex.RUnlock()

	default:
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}
