package main

import (
	"bytes"
	"fmt"
	"sync"
)

// best way to implement confinement is with lexical scoping.
// the following example shows a variable that needs to be used in
// multiple goroutines is created after, and passed into the goroutines.
func main() {

	printData := func(wg *sync.WaitGroup, data []byte) {
		defer wg.Done()
		var buff bytes.Buffer
		for _, b := range data {
			fmt.Fprintf(&buff, "%c", b)
		}
		fmt.Println(buff.String())
	}
	var wg sync.WaitGroup
	wg.Add(2)
	data := []byte("golang")
	go printData(&wg, data[:3])
	go printData(&wg, data[3:])
	wg.Wait()

}
