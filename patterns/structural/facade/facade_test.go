package facade_test

import (
	"fmt"
	"structural/facade"
	"testing"
)

// Test
func Test(t *testing.T) {
	// create facade object
	facade := facade.NewFacade()

	// use a simple method to do work that is a collection of underlying object methods
	facade.SetAllValues("hello, world")

	if facade.ObjOne().Val() != "hello, world" {
		t.Errorf("ValOne val is %s, and it should be hello, world", facade.ObjOne().Val())
	}
	if facade.ObjTwo().Val() != "hello, world" {
		t.Errorf("ValTwo val is %s, and it should be hello, world", facade.ObjTwo().Val())
	}
	if facade.ObjThree().Val() != "hello, world" {
		t.Errorf("ValThree val is %s, and it should be hello, world", facade.ObjThree().Val())
	}

}

// Example Test
func Example() {
	// create facade object
	facade := facade.NewFacade()

	// use a simple method to do work that is a collection of underlying object methods
	facade.SetAllValues("hello, world")

	fmt.Printf("ObjOne val is %s\n", facade.ObjOne().Val())
	fmt.Printf("ObjTwo val is %s\n", facade.ObjTwo().Val())
	fmt.Printf("ObjThree val is %s\n", facade.ObjThree().Val())

	//Output:ObjOne val is hello, world
	//ObjTwo val is hello, world
	//ObjThree val is hello, world
}
