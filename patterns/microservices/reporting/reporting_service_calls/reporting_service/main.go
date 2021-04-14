package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

// Service receives get from user and and retrieves data from required services.
// It caches the response locally to speed up the next request time
type Handler struct {
	cachedReports map[string]string
	mutex         *sync.RWMutex
}

func main() {

	h := &Handler{
		mutex:         &sync.RWMutex{},
		cachedReports: make(map[string]string, 0),
	}

	http.ListenAndServe(":9000", h)

}

// Serve HTTP Function to Implement RESTfull API
func (h *Handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:

		report1 := h.GetDataFromService("9001")
		report2 := h.GetDataFromService("9002")

		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(report1 + report2))

	default:
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *Handler) GetDataFromService(port string) string {

	// first check if report is cached locally
	h.mutex.RLock()
	cachedReport := h.cachedReports[port]
	h.mutex.RUnlock()
	if cachedReport != "" {
		return cachedReport
	}

	// if not, go and get it
	newReader1 := strings.NewReader("Request Report")

	resp, err := http.Post(fmt.Sprintf("http://localhost:%s", port), "application/text", newReader1)
	if err != nil {
		fmt.Printf("Error on post to %s: %s\n", port, err)
		return err.Error()
	}

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error on read of response from %s: %s\n", "9000", err)
		return err.Error()
	}

	h.mutex.Lock()
	h.cachedReports[port] = string(buf)
	h.mutex.Unlock()

	return string(buf)

}
