package state_test

import (
	"behavioral/state"
	"fmt"
	"testing"
)

func TestState(t *testing.T) {

	fan := state.NewFan()

	if fan.CurrentState().DebugName() != "state1" {
		t.Errorf("Fan is not in the correct state. Its state is %v, and it should be %v\n", fan.CurrentState().DebugName(), "state1")
	}
	fan.PullChain()
	if fan.CurrentState().DebugName() != "state2" {
		t.Errorf("Fan is not in the correct state. Its state is %v, and it should be %v\n", fan.CurrentState().DebugName(), "state2")
	}
	fan.PullChain()
	if fan.CurrentState().DebugName() != "state3" {
		t.Errorf("Fan is not in the correct state. Its state is %v, and it should be %v\n", fan.CurrentState().DebugName(), "state3")
	}
	fan.PullChain()
	if fan.CurrentState().DebugName() != "state1" {
		t.Errorf("Fan is not in the correct state. Its state is %v, and it should be %v\n", fan.CurrentState().DebugName(), "state1")
	}

}

func BenchmarkState(b *testing.B) {

	fan := state.NewFan()

	for i := 0; i < 1000000; i++ {
		fan.PullChain()
	}

}
func Example() {

	// create object with a state
	fan := state.NewFan()

	// use function to change state in object
	fan.PullChain()

	fmt.Println(fan.CurrentState().DebugName())
	//Output:state2

}
