package cor

type RequestType int

const (
	CONFERENCE = iota
	PURCHASE
)

type Request struct {
	requestType     RequestType
	amount          float64
	requestResponse string
}

// Getter method for the field requestType of type RequestType in the object Request
func (r *Request) RequestType() RequestType {
	return r.requestType
}

// Getter method for the field amount of type float64 in the object Request
func (r *Request) Amount() float64 {
	return r.amount
}

// Constructor for Request
func NewRequest(requestType RequestType, amount float64) *Request {
	o := new(Request)
	o.requestType = requestType
	o.amount = amount
	return o
}

// Getter method for the field requestResponse of type string in the object Request
func (r *Request) RequestResponse() string {
	return r.requestResponse
}

// Setter method for the field requestResponse of type string in the object Request
func (r *Request) SetRequestResponse(requestResponse string) {
	r.requestResponse = requestResponse
}
