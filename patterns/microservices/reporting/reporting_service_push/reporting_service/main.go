package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

type Handler struct {
	cachedReports map[string]string
	mutex         *sync.RWMutex
}

func main() {

	h := &Handler{
		cachedReports: make(map[string]string),
		mutex:         &sync.RWMutex{},
	}

	http.ListenAndServe(":9000", h)

}

// Serve HTTP Function to Implement RESTfull API
func (h *Handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:

		reportString := ""

		h.mutex.RLock()
		for port, report := range h.cachedReports {
			reportString += port + ": " + report + "\n"
		}
		h.mutex.RUnlock()

		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(reportString))

	case http.MethodPost:

		buf, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("Error on read of request: %s\n", err)
		}

		strArr := strings.Split(string(buf), ":")

		h.mutex.Lock()
		h.cachedReports[strArr[0]] = strArr[1]
		h.mutex.Unlock()

	default:
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}
