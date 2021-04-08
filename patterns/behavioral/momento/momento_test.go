package momento_test

import (
	"behavioral/momento"
	"fmt"
	"testing"
)

func TestMomento(t *testing.T) {

	caretaker := momento.NewCaretaker()

	employee := momento.NewEmployee(1, "Nick")

	caretaker.Save(employee)
	employee.SetName("Not Nick")
	if employee.Name() != "Not Nick" {
		t.Errorf("employee name did not get set correctly. It is %v, and it should be %v", employee.Name(), "Not Nick")
	}
	caretaker.Revert(employee)
	if employee.Name() != "Nick" {
		t.Errorf("employee name did not get set correctly. It is %v, and it should be %v", employee.Name(), "Nick")
	}

}

func BenchmarkMomento(b *testing.B) {

	caretaker := momento.NewCaretaker()

	employee := momento.NewEmployee(1, "Nick")

	for i := 0; i < 100000; i++ {
		caretaker.Save(employee)
		employee.SetName(fmt.Sprintf("%v", i))
	}
	for i := 0; i < 100000; i++ {
		caretaker.Revert(employee)
	}
	if employee.Name() != "Nick" {
		b.Errorf("employee name did not get set correctly. It is %v, and it should be %v", employee.Name(), "Nick")
	}
}

func Example() {

	// create object to hold saved states in
	caretaker := momento.NewCaretaker()

	// create object we will change
	employee := momento.NewEmployee(1, "Nick")

	// save state
	caretaker.Save(employee)

	// change state
	employee.SetName("Not Nick")

	// revert state to last save
	caretaker.Revert(employee)

	fmt.Printf("Employee Name is %v", employee.Name())

	// Output:Employee Name is Nick

}
