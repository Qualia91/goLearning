package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

func main() {

	wg := sync.WaitGroup{}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg.Add(1)
	go func() {
		defer wg.Done()

		if err := doSomething(ctx); err != nil {
			fmt.Println("Something error")
			cancel()
		}
	}()
}

func doSomething(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return errors.New("Done happend")
	}
	return errors.New("Error")
}
