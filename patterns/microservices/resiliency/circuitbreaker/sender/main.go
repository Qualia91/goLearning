package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"
)

// This handler speaks to 2 services. One will return straight away, one will hang forever.
// This handler will employ a timeout to cancel the read of the hanging service to save the application.
// The call is wrapped in a circuit breaker object which will keep track of how many times the hanging service fails.
// after 4 fails, it will flip and no calls will be able to be made to that anymore. Instead a default error is thrown.
type Handler struct {
}

const (
	ON = iota
	OFF
)

type CircuitBreaker struct {
	state               int
	lock                *sync.RWMutex
	breakCount          int
	breakLimit          int
	circuitbreakerError error
}

var cbOne CircuitBreaker
var cbTwo CircuitBreaker

func main() {

	// make circuitBreakers
	cbOne = CircuitBreaker{
		state:               ON,
		lock:                &sync.RWMutex{},
		breakCount:          0,
		breakLimit:          4,
		circuitbreakerError: errors.New("circuit break has flipped on service one"),
	}

	// make circuitBreakers
	cbTwo = CircuitBreaker{
		state:               ON,
		lock:                &sync.RWMutex{},
		breakCount:          0,
		breakLimit:          4,
		circuitbreakerError: errors.New("circuit break has flipped on service two"),
	}

	ah := new(Handler)

	http.ListenAndServe(":9000", ah)

}

// Serve HTTP Function to Implement RESTfull API
func (h Handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		wrapInCircuitBreaker(&cbOne, rw, "9001")
		wrapInCircuitBreaker(&cbTwo, rw, "9002")

	default:
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func wrapInCircuitBreaker(cb *CircuitBreaker, rw http.ResponseWriter, port string) {
	cb.lock.RLock()
	defer cb.lock.RUnlock()

	if cb.state == ON {

		if err := callServiceWithTimeout("Resquest", port, rw); err != nil {
			fmt.Printf("Timeout occurred waiting for service on port %s\n", port)
			cb.breakCount++
			if cb.breakCount > cb.breakLimit {
				cb.state = OFF
				fmt.Println("Circuit breaker for service one has flipped")
			}
		}

	} else {
		fmt.Println(cbOne.circuitbreakerError)
	}
}

func callServiceWithTimeout(requestText, port string, rw http.ResponseWriter) error {
	// create channel to manage timeout
	channel := make(chan []byte, 1)
	defer close(channel)

	// start send and receive in go func and pass return string to channel
	go func(channel chan []byte) {
		newReader := strings.NewReader(requestText)
		resp, err := http.Post(fmt.Sprintf("http://localhost:%s", port), "application/text", newReader)
		if err != nil {
			fmt.Printf("Error on post to %s: %s\n", port, err)
			return
		}

		buf, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error on read of response from %s: %s\n", port, err)
			return
		}

		channel <- buf
	}(channel)

	// use select statement on channel to impl actual timeout
	select {
	case <-channel:
		return nil
	case <-time.After(1 * time.Second):
		return fmt.Errorf("service on port %s did not return in time", port)
	}

}
