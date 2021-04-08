package mediator

type Mediator interface {
	Register(obj *InteractableObject)
}

type InteractableObject struct {
	state bool
}

type LightMediator struct {
	registry []*InteractableObject
}

// Implements Mediator
func (lightMediator *LightMediator) Register(io *InteractableObject) {
	lightMediator.registry = append(lightMediator.registry, io)
}

// functionality to be applied to all registered objects
func (lm *LightMediator) ToggleState() {
	for index, intObj := range lm.registry {
		lm.registry[index].SetState(!intObj.state)
	}
}

// Constructor for InteractableObject
func NewInteractableObject(state bool) *InteractableObject {
	o := new(InteractableObject)
	o.state = state
	return o
}

// Constructor for LightMediator
func NewLightMediator() *LightMediator {
	o := new(LightMediator)
	o.registry = make([]*InteractableObject, 0)
	return o
}

// Getter method for the field state of type bool in the object InteractableObject
func (i *InteractableObject) State() bool {
	return i.state
}

// Setter method for the field state of type bool in the object InteractableObject
func (i *InteractableObject) SetState(state bool) {
	i.state = state
}
