package builder_test

import (
	"creational/builder"
	"fmt"
	"testing"
)

// Test
func Test(t *testing.T) {

	builder := builder.NewBuilder()
	builtObj := builder.VarOne("1").VarTwo("2").VarThree("3").VarFour("4").Build()

	if builtObj.VarOne() != "1" {
		t.Errorf("VarOne was %v, and was supposed to be 1", builtObj.VarOne())
	}
	if builtObj.VarTwo() != "2" {
		t.Errorf("VarTwo was %v, and was supposed to be 2", builtObj.VarTwo())
	}
	if builtObj.VarThree() != "3" {
		t.Errorf("VarThree was %v, and was supposed to be 3", builtObj.VarThree())
	}
	if builtObj.VarFour() != "4" {
		t.Errorf("VarFour was %v, and was supposed to be 4", builtObj.VarFour())
	}

}

// Example Test
func Example() {

	// make builder
	builder := builder.NewBuilder()

	// edit value and build object
	builtObj := builder.VarOne("1").Build()

	fmt.Printf("VarOne was %v", builtObj.VarOne())

	//Output:VarOne was 1
}
