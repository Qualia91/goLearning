package cor

type Handler interface {
	SetSuccessor(Handler)
	HandleRequest(Request) string
}

type HandleRequest func(Request, Handler) string

type HandlerImpl struct {
	successor     Handler
	handleRequest HandleRequest
}

// Constructor for HandlerImpl
func NewHandlerImpl(handleRequest HandleRequest) *HandlerImpl {
	o := new(HandlerImpl)
	o.handleRequest = handleRequest
	return o
}

// Setter method for the field successor of type Handler in the object HandlerImpl
func (h *HandlerImpl) SetSuccessor(successor Handler) {
	h.successor = successor
}

func (h *HandlerImpl) HandleRequest(request Request) string {
	return h.handleRequest(request, h.successor)
}
