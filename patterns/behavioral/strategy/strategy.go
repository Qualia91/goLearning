package strategy

type ObjectToValidateInterface interface {
	Value() int
}

type ObjectToValidate struct {
	name         string
	value        int
	strategyFunc func(obj ObjectToValidateInterface) bool
}

// Constructor for ObjectToValidate
func NewObjectToValidate(stratFunc func(obj ObjectToValidateInterface) bool) *ObjectToValidate {
	o := new(ObjectToValidate)
	o.strategyFunc = stratFunc
	return o
}

// Getter method for the field name of type string in the object ObjectToValidate
func (o *ObjectToValidate) Name() string {
	return o.name
}

// Setter method for the field name of type string in the object ObjectToValidate
func (o *ObjectToValidate) SetName(name string) {
	o.name = name
}

// Getter method for the field value of type int in the object ObjectToValidate
func (o *ObjectToValidate) Value() int {
	return o.value
}

// Setter method for the field value of type int in the object ObjectToValidate
func (o *ObjectToValidate) SetValue(value int) {
	o.value = value
}

// Getter method for the field valid of type bool in the object ObjectToValidate
func (o *ObjectToValidate) Valid() bool {
	return o.strategyFunc(o)
}
