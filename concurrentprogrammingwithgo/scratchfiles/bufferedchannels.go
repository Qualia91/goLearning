package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	// seconds param is buffer
	ch := make(chan int, 1)

	wg.Add(2)

	go func(ch chan int, wg *sync.WaitGroup) {
		// received from channel
		fmt.Println(<-ch)
		wg.Done()
	}(ch, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		// send to channel
		ch <- 42
		ch <- 27
		wg.Done()
	}(ch, wg)

	wg.Wait()
}
