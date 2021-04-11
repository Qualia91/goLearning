/*Facade pattern.

Will use the Facade object with a few methods in to call a collection
of methods in objects within Facade that collectively do something
*/

package facade

type Facade struct {
	objOne   *objOne
	objTwo   *objTwo
	objThree *objThree
}

type objOne struct {
	val string
}

type objTwo struct {
	val string
}

type objThree struct {
	val string
}

// Constructor for Facade
func NewFacade() *Facade {
	o := new(Facade)
	o.objOne = new(objOne)
	o.objTwo = new(objTwo)
	o.objThree = new(objThree)
	return o
}

func (f *Facade) SetAllValues(val string) {
	f.objOne.val = val
	f.objTwo.val = val
	f.objThree.val = val
}

// Getter method for the field objOne of type *objOne in the object Facade
func (f *Facade) ObjOne() *objOne {
	return f.objOne
}

// Getter method for the field objTwo of type *objTwo in the object Facade
func (f *Facade) ObjTwo() *objTwo {
	return f.objTwo
}

// Getter method for the field objThree of type *objThree in the object Facade
func (f *Facade) ObjThree() *objThree {
	return f.objThree
}

// Getter method for the field val of type string in the object objOne
func (o *objOne) Val() string {
	return o.val
}

// Getter method for the field val of type string in the object objTwo
func (o *objTwo) Val() string {
	return o.val
}

// Getter method for the field val of type string in the object objThree
func (o *objThree) Val() string {
	return o.val
}
