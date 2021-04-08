package iterator

type InteratorClosure interface {
	Next() interface{}
	HasNext() bool
}

type Collection interface {
	IterC() InteratorClosure
	Add(interface{})
}
