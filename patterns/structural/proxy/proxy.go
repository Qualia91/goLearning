package proxy

type MyObjectInterface interface {
	Apply(str string)
}

type MyObject struct {
	val string
}

type Proxy struct {
	mo *MyObject
}

// Constructor for MyObject
func NewMyObject(val string) *MyObject {
	o := new(MyObject)
	o.val = val
	return o
}

// Getter method for the field val of type string in the object MyObject
func (m *MyObject) Val() string {
	return m.val
}

// Constructor for Proxy
func NewProxy(mo *MyObject) *Proxy {
	o := new(Proxy)
	o.mo = mo
	return o
}

// Getter method for the field mo of type *MyObject in the object Proxy
func (p *Proxy) Val() string {
	return p.mo.val + " as found by the proxy"
}

// Implements MyObjectInterface
func (myObject *MyObject) Apply(str string) {
	myObject.val = str
}

// Implements MyObjectInterface
func (proxy *Proxy) Apply(str string) {
	proxy.mo.val = "Proxy set val " + str
}
