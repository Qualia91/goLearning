package main

import (
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int)

	wg.Add(2)

	// receive only channel
	go func(ch <-chan int, wg *sync.WaitGroup) {
		for msg := range ch {
			println(msg)
		}
		wg.Done()

	}(ch, wg)

	// send only channel
	go func(ch chan<- int, wg *sync.WaitGroup) {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
	}(ch, wg)

	wg.Wait()

}
