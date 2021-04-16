package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	done := make(chan interface{})
	intStream := Generator(done, 1, 2, "Hello", 4, false, 6)
	for v := range FanIn(done, FanOut(done, Curry(StageOne, 4), intStream, 10)) {
		fmt.Println(v)
	}

	// currying example
	for v := range Curry(StageOne, 4)(done, intStream) {
		fmt.Println(v)
	}

}

// function to Curry the stage function into a form that can be used in FanOut
func Curry(origFn func(<-chan interface{}, <-chan interface{}, int) <-chan interface{}, val int) func(<-chan interface{}, <-chan interface{}) <-chan interface{} {
	return func(done <-chan interface{}, stream <-chan interface{}) <-chan interface{} {
		return origFn(done, stream, val)
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

// fan in stage
func FanIn(done <-chan interface{}, streams []<-chan interface{}) <-chan interface{} {

	// create waitgroup to wait for all channels to be done
	wg := &sync.WaitGroup{}

	multiplexStream := make(chan interface{})

	// create function that takes in channel and adds read from channel to multipleStream
	multiplex := func(c <-chan interface{}) {
		defer wg.Done()
		for i := range c {
			select {
			case <-done:
				return
			case multiplexStream <- i:
			}
		}
	}

	// add to wg
	wg.Add(len(streams))

	// iterate over streams and multiple
	for _, c := range streams {
		go multiplex(c)
	}

	// create a goroutine to close multiplex stream when all streams are done
	// this is done in a goroutine so the multiplestream is returned straight away to
	// stages after, as we aren't doing any work here, just creating chains of streams (pipelines)
	go func() {
		wg.Wait()
		close(multiplexStream)
	}()

	return multiplexStream
}

// fan out stage
func FanOut(done <-chan interface{}, stage func(<-chan interface{}, <-chan interface{}) <-chan interface{}, stream <-chan interface{}, fanTo int) []<-chan interface{} {

	stages := make([]<-chan interface{}, fanTo)
	for i := 0; i < fanTo; i++ {
		stages[i] = stage(done, stream)
	}

	return stages
}

// Example of a stage that takes a while to do work
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
	time.Sleep(5 * time.Second)
	switch inputVal.(type) {
	case int:
		return operationVal * inputVal.(int)
	default:
		return inputVal
	}
}
