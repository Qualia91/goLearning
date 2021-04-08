package mediator_test

import (
	"behavioral/mediator"
	"fmt"
	"testing"
)

func TestMediator(t *testing.T) {

	lm := mediator.NewLightMediator()

	intObj1 := mediator.NewInteractableObject(true)
	intObj2 := mediator.NewInteractableObject(false)

	lm.Register(intObj1)
	lm.Register(intObj2)

	lm.ToggleState()

	if intObj1.State() != false {
		t.Errorf("intObj1's state is %t, while it should be %t", intObj1.State(), false)
	}

	if intObj2.State() != true {
		t.Errorf("intObj2's state is %t, while it should be %t", intObj2.State(), true)
	}

}

func BenchmarkMediator(b *testing.B) {

	count := 100000
	objs := make([]*mediator.InteractableObject, count)

	lm := mediator.NewLightMediator()

	for i := 0; i < count; i++ {
		intObj := mediator.NewInteractableObject(true)
		lm.Register(intObj)
		objs[i] = intObj
	}

	lm.ToggleState()

	for index, obj := range objs {
		if obj.State() != false {
			b.Errorf("obj number %v state is %t, while it should be %t", index, obj.State(), false)
		}
	}

}

func Example() {

	lm := mediator.NewLightMediator()

	intObj1 := mediator.NewInteractableObject(true)
	intObj2 := mediator.NewInteractableObject(false)

	lm.Register(intObj1)
	lm.Register(intObj2)

	lm.ToggleState()

	fmt.Printf("State of intObj1 is %t, and state of intObj2 is %t", intObj1.State(), intObj2.State())

	// Output:State of intObj1 is false, and state of intObj2 is true

}
