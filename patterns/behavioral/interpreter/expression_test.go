package interpreter_test

import (
	"behavioral/interpreter"
	"fmt"
	"testing"
)

// helper function to build simple interpreter tree
func buildInterpreterTree() interpreter.Expression {
	interp1 := interpreter.NewTerminalExpression("Lions")
	interp2 := interpreter.NewTerminalExpression("Tigers")
	interp3 := interpreter.NewTerminalExpression("Bears")

	// tigers and bears
	interp4 := interpreter.NewAndExpression(interp2, interp3)

	// lions or (tigers and bears)
	interp5 := interpreter.NewOrExpression(interp1, interp4)

	// bears and (lions or (tigers and bears))
	return interpreter.NewAndExpression(interp3, interp5)
}

func TestInterpreter(t *testing.T) {

	define := buildInterpreterTree()

	testCases := []struct {
		context  string
		expected bool
	}{
		{context: "a", expected: false},
		{context: "Tigers", expected: false},
		{context: "Bears", expected: false},
		{context: "Lions Tigers", expected: false},
		{context: "Lions Bears", expected: true},
		{context: "Tigers Bears", expected: true},
	}
	for _, tC := range testCases {
		t.Run(tC.context, func(t *testing.T) {
			if define.Interpret(tC.context) != tC.expected {
				t.Errorf("%v was expected to return %t, but instead returned %t", tC.context, tC.expected, !tC.expected)
			}
		})
	}
}

func BenchmarkInterpreter(b *testing.B) {

	b.StartTimer()

	define := buildInterpreterTree()

	testCases := []struct {
		context  string
		expected bool
	}{
		{context: "a", expected: false},
		{context: "Tigers", expected: false},
		{context: "Bears", expected: false},
		{context: "Lions Tigers", expected: false},
		{context: "Lions Bears", expected: true},
		{context: "Tigers Bears", expected: true},
	}
	for _, tC := range testCases {
		b.Run(tC.context, func(b *testing.B) {
			if define.Interpret(tC.context) != tC.expected {
				b.Errorf("%v was expected to return %t, but instead returned %t", tC.context, tC.expected, !tC.expected)
			}
		})
	}
}

func Example() {
	// using interpreter
	context := "Tigers Bears"

	interp1 := interpreter.NewTerminalExpression("Lions")
	interp2 := interpreter.NewTerminalExpression("Tigers")
	interp3 := interpreter.NewTerminalExpression("Bears")

	// tigers and bears
	interp4 := interpreter.NewAndExpression(interp2, interp3)

	// lions or (tigers and bears)
	interp5 := interpreter.NewOrExpression(interp1, interp4)

	// bears and (lions or (tigers and bears))
	interp6 := interpreter.NewAndExpression(interp3, interp5)

	fmt.Printf("%v is %t\n", context, interp6.Interpret(context))

	// Output: Tigers Bears is true

}
