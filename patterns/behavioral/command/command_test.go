package command_test

import (
	"behavioral/command"
	"fmt"
	"testing"
)

func BenchmarkCommand(b *testing.B) {
	receiver := command.NewReceiver()
	invoker := command.NewInvoker()

	onCommand := command.NewCommand(
		func() {
			receiver.On = true
		},
		func() {
			receiver.On = false
		},
	)

	offCommand := command.NewCommand(
		func() {
			receiver.On = false
		},
		func() {
			receiver.On = true
		},
	)

	for i := 0; i < 100000; i++ {
		invoker.Execute(*onCommand)
		invoker.Execute(*offCommand)
	}

	for i := 0; i < 100000; i++ {
		invoker.Undo()
		invoker.Undo()
	}

}

func TestCommand(t *testing.T) {
	receiver := command.NewReceiver()
	invoker := command.NewInvoker()

	onCommand := command.NewCommand(
		func() {
			receiver.On = true
		},
		func() {
			receiver.On = false
		},
	)

	offCommand := command.NewCommand(
		func() {
			receiver.On = false
		},
		func() {
			receiver.On = true
		},
	)

	if receiver.On {
		t.Errorf("Receivers value should be %t, but is %t", !receiver.On, receiver.On)
	}
	invoker.Execute(*onCommand)
	if !receiver.On {
		t.Errorf("Receivers value should be %t, but is %t", !receiver.On, receiver.On)
	}
	invoker.Execute(*offCommand)
	if receiver.On {
		t.Errorf("Receivers value should be %t, but is %t", !receiver.On, receiver.On)
	}
	invoker.Undo()
	if !receiver.On {
		t.Errorf("Receivers value should be %t, but is %t", !receiver.On, receiver.On)
	}
}

func Example() {

	// create receiver (thing being updated by command)
	receiver := command.NewReceiver()

	// create invoker (were the commands are executed and stored)
	invoker := command.NewInvoker()

	// create commands
	onCommand := command.NewCommand(
		func() {
			receiver.On = true
		},
		func() {
			receiver.On = false
		},
	)

	invoker.Execute(*onCommand)

	invoker.Undo()

	fmt.Printf("Invoker on state is %t", receiver.On)

	// Output: Invoker on state is false
}
