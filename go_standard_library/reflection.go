package main

import (
	"fmt"
	"reflect"
	"runtime"
	"time"

	"example.com/test/media"
)

func main() {
	fmt.Println()

	myFav := media.Movie{}
	myFav.NewMovie("Starship Trooper", media.R, 110.0)

	myFavPtr := &media.Movie{}
	myFavPtr.NewMovie("Starship Trooper", media.R, 110.0)

	fmt.Printf("My fave mov is %s, %v, %f\n", myFav.GetTitle(), myFav.GetRating(), myFav.GetBoxOffice())

	myFav.SetTitle("Still STarship Troopers")

	fmt.Printf("My fave mov is %s, %v, %f\n", myFav.GetTitle(), myFav.GetRating(), myFav.GetBoxOffice())

	fmt.Printf("Type: %v\n", reflect.TypeOf(myFav))
	fmt.Printf("Value: %v\n", reflect.ValueOf(myFav))
	fmt.Printf("Kind: %v\n", reflect.ValueOf(myFav).Kind())

	timedFunc := MakeTimedFunction(Printer).(func(string) string)
	timedFuncTwo := MakeTimedFunction(typeCheck).(func(interface{}) bool)
	timedFunc(myFav.GetTitle())
	timedFuncTwo(myFav)
	typeCheck(myFav)
	typeCheck(myFavPtr)

}

func Printer(str string) string {
	return str
}

func MakeTimedFunction(f interface{}) interface{} {
	rf := reflect.TypeOf(f)

	if rf.Kind() != reflect.Func {
		panic("Not a function")
	}

	vf := reflect.ValueOf(f)
	wrapperF := reflect.MakeFunc(rf, func(in []reflect.Value) []reflect.Value {
		start := time.Now()
		out := vf.Call(in)
		end := time.Now()
		fmt.Printf("calling %s took %v\n", runtime.FuncForPC(vf.Pointer()).Name(), end.Sub(start))
		return out
	})

	return wrapperF.Interface()

}

func typeCheck(p interface{}) bool {
	if reflect.ValueOf(p).Kind() == reflect.Struct {
		fmt.Println("Struct")
	} else if reflect.ValueOf(p).Kind() == reflect.Ptr {
		fmt.Println("Pointer")
	} else {
		fmt.Println("Type not found")
	}
	return true
}
