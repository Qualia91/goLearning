package adapter

type TargetModel struct {
	x string
}

// Implements ObjInt
func (targetModel *TargetModel) WorkFunction() string {
	return "Target model is now doing work"
}

// Constructor for TargetModel
func NewTargetModel(om OriginalModel) *TargetModel {
	o := new(TargetModel)
	o.x = om.a + om.b + om.c
	return o
}

// Getter method for the field x of type string in the object TargetModel
func (t *TargetModel) X() string {
	return t.x
}

// Setter method for the field x of type string in the object TargetModel
func (t *TargetModel) SetX(x string) {
	t.x = x
}
