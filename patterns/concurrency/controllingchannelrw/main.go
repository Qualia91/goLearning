package main

import (
	"fmt"
	"sync"
)

// this demonstrates how passing in a read only or write only version of
// a channel can help to maintain r/w of channels within go routines they are passed to,
// and display to users of functions using type signatures.
// The channel owner is the main go routine and that is the only one that has access to the
// read and write of the channel
func main() {

	// create wait group to stop prog exiting prematurely
	wg := &sync.WaitGroup{}

	// create un-buffered channel
	ch := make(chan string)

	// create a write only go routine
	wg.Add(1)
	go func(ch chan<- string) {
		defer wg.Done()

		// aside: As this channel's purpose is to pass one message from this goroutine to another go routine,
		// it makes sense to close the channel here as this go routine know exactly when then channel is no
		// longer in use
		defer close(ch)

		// this line shows error in ide
		// <-ch
		ch <- "hello"
	}(ch)

	// create a read only go routine
	wg.Add(1)
	go func(ch <-chan string) {
		defer wg.Done()
		// this line shows error in ide
		// ch <- "hello"
		fmt.Println(<-ch)
	}(ch)

	wg.Wait()
}
