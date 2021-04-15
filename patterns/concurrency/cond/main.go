package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	// create a wait group to make sure everything finishes before end of program
	wg := sync.WaitGroup{}

	// create a cond
	c := sync.NewCond(&sync.Mutex{})

	// create 10 goroutines that will all wait on a
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(num int) {
			c.L.Lock()
			defer func() {
				c.L.Unlock()
				wg.Done()
			}()
			c.Wait()
			fmt.Printf("Go Routine %v just ran\n", num)
		}(i)
	}

	// just ignore this race condition...
	time.Sleep(5 * time.Second)

	c.Broadcast()
	wg.Wait()
}
