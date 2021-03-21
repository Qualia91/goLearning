package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
)

func main() {

	fmt.Println("Current version " + runtime.Version())

	args := os.Args

	for i := 0; i < len(args); i++ {
		println(args[i])
	}

	args_slice := args[1:]

	if len(args_slice) != 2 {
		fmt.Println("You must enter args")
	} else {
		mealTot, _ := strconv.ParseFloat(args_slice[0], 32)
		tipAmount, _ := strconv.ParseFloat(args_slice[1], 32)
		fmt.Println(mealTot)
		fmt.Println(tipAmount)
		fmt.Printf("Your meal tot will be %.2f\n", calculateTot(float32(mealTot), float32(tipAmount)))
	}

}

func calculateTot(mealTot float32, tipAmount float32) float32 {
	return mealTot + (mealTot * (tipAmount / 100))
}
