package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan string)

	go func(channel chan<- string) {
		defer close(channel)
		time.Sleep(1 * time.Second)
		channel <- "Work is done"
	}(ch)

	for {
		select {
		case str, isClose := <-ch:
			fmt.Println(str)
			if isClose {
				fmt.Println("Channel closed")
				return
			}

		default:
			fmt.Println("Waiting and doing other work")
		}
	}

}
