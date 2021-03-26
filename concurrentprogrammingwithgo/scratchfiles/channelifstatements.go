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
		if msg, ok := <-ch; ok {
			fmt.Println(msg)
		} else {
			fmt.Println("Channel is closed")
		}
		wg.Done()

	}(ch, wg)

	// send only channel
	go func(ch chan<- int, wg *sync.WaitGroup) {
		// this closes a channel
		close(ch)
		wg.Done()
	}(ch, wg)

	wg.Wait()

}
