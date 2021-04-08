package cor_test

import (
	"behavioral/cor"
	"fmt"
	"testing"
)

func TestCor(t *testing.T) {
	req10 := cor.NewRequest(cor.RequestType(cor.CONFERENCE), 10)
	req1000 := cor.NewRequest(cor.RequestType(cor.CONFERENCE), 1000)

	cor1HandleRequest := func(request cor.Request, successor cor.Handler) string {
		if request.Amount() < 100 {
			return fmt.Sprintf("Amount is %v so cor1 can handle\n", request.Amount())
		}
		return successor.HandleRequest(request)
	}
	cor1 := cor.NewHandlerImpl(cor1HandleRequest)

	cor2HandleRequest := func(request cor.Request, successor cor.Handler) string {
		return fmt.Sprintf("Amount is %v so cor2 can handle\n", request.Amount())
	}
	cor2 := cor.NewHandlerImpl(cor2HandleRequest)

	cor1.SetSuccessor(cor2)

	req10Return := cor1.HandleRequest(*req10)
	req1000Return := cor1.HandleRequest(*req1000)

	if req10Return != "Amount is 10 so cor1 can handle\n" {
		t.Errorf("Req10 returned string %v, instead of expected string of \"Amount is 10 so cor1 can handle\"", req10Return)
	}

	if req1000Return != "Amount is 1000 so cor2 can handle\n" {
		t.Errorf("Req1000 returned string %v, instead of expected string of \"Amount is 1000 so cor2 can handle\"", req1000Return)
	}

}

func BenchmarkCorStress(b *testing.B) {

	b.StartTimer()

	req1000 := cor.NewRequest(cor.RequestType(cor.CONFERENCE), 1000)

	var prev *cor.HandlerImpl
	var first *cor.HandlerImpl
	for i := 0; i < 1000000; i++ {
		cor1HandleRequest := func(request cor.Request, successor cor.Handler) string {
			if request.Amount() > 20000000 {
				return fmt.Sprintf("Amount is %v so cor1 can handle\n", request.Amount())
			}
			if successor != nil {
				return successor.HandleRequest(request)
			}
			return "Can't Handle\n"
		}
		cor1 := cor.NewHandlerImpl(cor1HandleRequest)
		if first == nil {
			first = cor1
		}
		if prev != nil {
			prev.SetSuccessor(cor1)
		}
		prev = cor1
	}

	prev.HandleRequest(*req1000)
}

func Example() {
	// Create a request you want to go through the chain.
	req10 := cor.NewRequest(cor.RequestType(cor.CONFERENCE), 10)

	// Create handler 1's function
	cor1HandleRequest := func(request cor.Request, successor cor.Handler) string {
		if request.Amount() < 100 {
			return fmt.Sprintf("Amount is %v so cor1 can handle", request.Amount())
		}
		return successor.HandleRequest(request)
	}

	// create handler 1
	cor1 := cor.NewHandlerImpl(cor1HandleRequest)

	// create handler 2's functions
	cor2HandleRequest := func(request cor.Request, successor cor.Handler) string {
		return fmt.Sprintf("Amount is %v so cor2 can handle", request.Amount())
	}

	// create handler 2
	cor2 := cor.NewHandlerImpl(cor2HandleRequest)

	// create chain of handler 1 -> handler 2
	cor1.SetSuccessor(cor2)

	// run request through chain and print result
	fmt.Println(cor1.HandleRequest(*req10))

	// Output: Amount is 10 so cor1 can handle
}
