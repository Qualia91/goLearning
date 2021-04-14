package adapter_test

import (
	"fmt"
	"structural/adapter"
	"testing"
)

// this function uses the models to do work. It was initially designed to work with OriginalModel,
// but now uses the interface so it can use target model too.
func UseObjectAndWork(o adapter.ObjInt) string {
	return o.WorkFunction()
}

// Test
func Test(t *testing.T) {

	// create original model
	om := adapter.NewOriginalModel("A", "B", "C")

	// get return from working functions
	omw := UseObjectAndWork(om)

	// create target model
	tm := adapter.NewTargetModel(*om)

	// get return from working function
	tmw := UseObjectAndWork(tm)

	if omw != "Has done some work" {
		t.Errorf("Work from Original Object was %s, and it should have been \"Has done some work\"", omw)
	}

	if tmw != "Target model is now doing work" {
		t.Errorf("Work from Original Object was %s, and it should have been \"Target model is now doing work\"", tmw)
	}

}

// Example Test
func Example() {

	// create original model
	om := adapter.NewOriginalModel("A", "B", "C")

	// get return from working functions
	omw := UseObjectAndWork(om)

	// create target model
	tm := adapter.NewTargetModel(*om)

	// get return from working function
	tmw := UseObjectAndWork(tm)

	fmt.Println(omw)
	fmt.Println(tmw)

	//Output:Has done some work
	//Target model is now doing work

}
