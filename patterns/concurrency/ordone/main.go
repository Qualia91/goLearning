package main

import "fmt"

func main() {

	done := make(chan interface{})
	defer close(done)
	myChan := make(chan interface{}, 2)
	myChan <- "Hello"
	myChan <- "World"
	close(myChan)

	for val := range OrDone(done, myChan) {
		fmt.Println(val)
	}
}

func OrDone(done <-chan interface{}, c <-chan interface{}) <-chan interface{} {
	valStream := make(chan interface{})
	go func() {
		defer close(valStream)
		for {
			select {
			case <-done:
				return
			case v, ok := <-c:
				if !ok {
					return
				}
				select {
				case valStream <- v:
				case <-done:
				}
			}
		}
	}()
	return valStream
}
