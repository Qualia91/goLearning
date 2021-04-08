package iterator

type MyCollection struct {
	data []interface{}
}

type MyIteratorC struct {
	mc    *MyCollection
	index int
}

// Implements InteratorClosure
func (mi *MyIteratorC) Next() interface{} {
	nextData := mi.mc.data[mi.index]
	mi.index++
	return nextData
}

// Implements InteratorClosure
func (mi *MyIteratorC) HasNext() bool {
	return len(mi.mc.data) > mi.index
}

// Implements Collection
func (mc *MyCollection) IterC() InteratorClosure {
	return NewMyIteratorC(mc)
}

// Constructor for MyCollection
func NewMyCollection() *MyCollection {
	o := new(MyCollection)
	return o
}

// Constructor for MyIteratorC
func NewMyIteratorC(mc *MyCollection) *MyIteratorC {
	o := new(MyIteratorC)
	o.mc = mc
	o.index = 0
	return o
}

// Implements Collection
func (myCollection *MyCollection) Add(obj interface{}) {
	myCollection.data = append(myCollection.data, obj)
}
