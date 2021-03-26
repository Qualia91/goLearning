package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int)

	wg.Add(2)

	// receive only channel
	go func(ch <-chan int, wg *sync.WaitGroup) {
		// received from channel
		fmt.Println(<-ch)
		wg.Done()
	}(ch, wg)

	// send only channel
	go func(ch chan<- int, wg *sync.WaitGroup) {
		// send to channel
		ch <- 42
		wg.Done()
	}(ch, wg)

	wg.Wait()

	// this closes a channel
	close(ch)
}
