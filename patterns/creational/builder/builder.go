package builder

type Builder struct {
	varOne   string
	varTwo   string
	varThree string
	varFour  string
}

type BuiltObject struct {
	varOne   string
	varTwo   string
	varThree string
	varFour  string
}

// Constructor for Builder
func NewBuilder() *Builder {
	o := new(Builder)
	return o
}

func (b *Builder) Build() *BuiltObject {
	o := new(BuiltObject)
	o.varOne = b.varOne
	o.varTwo = b.varTwo
	o.varThree = b.varThree
	o.varFour = b.varFour
	return o
}

func (b *Builder) VarOne(s string) *Builder {
	b.varOne = s
	return b
}

func (b *Builder) VarTwo(s string) *Builder {
	b.varTwo = s
	return b
}

func (b *Builder) VarThree(s string) *Builder {
	b.varThree = s
	return b
}

func (b *Builder) VarFour(s string) *Builder {
	b.varFour = s
	return b
}

// Getter method for the field varOne of type string in the object BuiltObject
func (b *BuiltObject) VarOne() string {
	return b.varOne
}

// Getter method for the field varTwo of type string in the object BuiltObject
func (b *BuiltObject) VarTwo() string {
	return b.varTwo
}

// Getter method for the field varThree of type string in the object BuiltObject
func (b *BuiltObject) VarThree() string {
	return b.varThree
}

// Getter method for the field varFour of type string in the object BuiltObject
func (b *BuiltObject) VarFour() string {
	return b.varFour
}
