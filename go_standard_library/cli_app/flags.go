package main

import (
	"flag"
	"fmt"
	"runtime"
)

func main() {

	archPtr := flag.String("arch", "x86", "CPU Type")

	flag.Parse()

	switch *archPtr {
	case "x86":
		fmt.Println("Running 32 bit mode")
	case "AMD64":
		fmt.Println("Running in 64 bit mode")
	default:
		fmt.Println("Other")
	}

	fmt.Println("Current version " + runtime.Version())

}
