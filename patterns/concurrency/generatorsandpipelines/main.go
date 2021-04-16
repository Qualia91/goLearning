package main

import "fmt"

func main() {
	done := make(chan interface{})
	intStream := Generator(done, 1, 2, "Hello", 4, false, 6)
	for v := range StageOne(done, intStream, 4) {
		fmt.Println(v)
	}

	// example of handy generators
	for i := range Take(done, Repeat(done, func() int { return 100 }()), 5) {
		fmt.Println(i)
	}

}

// uses done channel to get rid of go routine leaks
func Generator(done <-chan interface{}, vals ...interface{}) <-chan interface{} {

	// make channel the length of incoming list (optimization)
	stream := make(chan interface{}, len(vals))
	go func() {
		defer close(stream)
		for _, i := range vals {
			select {
			case <-done:
				return
			case stream <- i:
			}
		}
	}()
	return stream

}

// Example of a stage
func StageOne(done <-chan interface{}, stream <-chan interface{}, operationVal int) <-chan interface{} {

	outputStream := make(chan interface{}, len(stream))
	go func() {
		defer close(outputStream)
		for i := range stream {
			select {
			case <-done:
				return
			case outputStream <- applyOperation(i, operationVal):
			}
		}
	}()
	return outputStream

}

func applyOperation(inputVal interface{}, operationVal int) interface{} {
	switch inputVal.(type) {
	case int:
		return operationVal * inputVal.(int)
	default:
		return inputVal
	}
}
