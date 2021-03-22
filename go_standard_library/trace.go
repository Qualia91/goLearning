package main

import (
	"log"
	"math/rand"
	"os"
	"runtime/trace"
)

func main() {

	f, err := os.Create("Trace.out")
	if err != nil {
		log.Fatalf("We did not create trace file! %v\n", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("Failed to close trace %v\n", err)
		}
	}()

	if err := trace.Start(f); err != nil {
		log.Fatalf("We failed to start trace: %d\n", err)
	}
	defer trace.Stop()

	AddRandomRumbers()

}

func AddRandomRumbers() int {
	firstNumber := rand.Intn(100)
	secondNumber := rand.Intn(100)
	return firstNumber + secondNumber
}
