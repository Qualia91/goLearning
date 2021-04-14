package adapter

type OriginalModel struct {
	a string
	b string
	c string
}

// Constructor for OriginalModel
func NewOriginalModel(a string, b string, c string) *OriginalModel {
	o := new(OriginalModel)
	o.a = a
	o.b = b
	o.c = c
	return o
}

func (o *OriginalModel) WorkFunction() string {
	return "Has done some work"
}

// Getter method for the field a of type string in the object OriginalModel
func (o *OriginalModel) A() string {
	return o.a
}

// Setter method for the field a of type string in the object OriginalModel
func (o *OriginalModel) SetA(a string) {
	o.a = a
}

// Getter method for the field b of type string in the object OriginalModel
func (o *OriginalModel) B() string {
	return o.b
}

// Setter method for the field b of type string in the object OriginalModel
func (o *OriginalModel) SetB(b string) {
	o.b = b
}

// Getter method for the field c of type string in the object OriginalModel
func (o *OriginalModel) C() string {
	return o.c
}

// Setter method for the field c of type string in the object OriginalModel
func (o *OriginalModel) SetC(c string) {
	o.c = c
}
