package strategy_test

import (
	"behavioral/strategy"
	"fmt"
	"testing"
)

// Test
func Test(t *testing.T) {

	validationStratOne := func(obj strategy.ObjectToValidateInterface) bool {
		return obj.Value() < 0
	}

	validationStratTwo := func(obj strategy.ObjectToValidateInterface) bool {
		return obj.Value() > 0
	}

	verOne := strategy.NewObjectToValidate(validationStratOne)
	verTwo := strategy.NewObjectToValidate(validationStratTwo)

	verOne.SetValue(1)
	verTwo.SetValue(1)

	if verOne.Valid() {
		t.Errorf("verOne valid is %t, and it should be false", verOne.Valid())
	}

	if !verTwo.Valid() {
		t.Errorf("verTwo valid is %t, and it should be true", verTwo.Valid())
	}

}

// Example Test
func Example() {

	// create 2 different validation functions
	validationStratOne := func(obj strategy.ObjectToValidateInterface) bool {
		return obj.Value() < 0
	}

	validationStratTwo := func(obj strategy.ObjectToValidateInterface) bool {
		return obj.Value() > 0
	}

	// create 2 objects with the validation functions
	verOne := strategy.NewObjectToValidate(validationStratOne)
	verTwo := strategy.NewObjectToValidate(validationStratTwo)

	// set values in them
	verOne.SetValue(1)
	verTwo.SetValue(1)

	fmt.Printf("verOne valid is %t\n", verOne.Valid())
	fmt.Printf("verTwo valid is %t\n", verTwo.Valid())

	//Output:verOne valid is false
	//verTwo valid is true
}
