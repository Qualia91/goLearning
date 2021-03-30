package cor_test

import (
	"behavioral/cor"
	"fmt"
	"testing"
)

func TestCor(t *testing.T) {
	req10 := cor.NewRequest(cor.RequestType(cor.CONFERENCE), 10)
	req1000 := cor.NewRequest(cor.RequestType(cor.CONFERENCE), 1000)

	cor1HandleRequest := func(request cor.Request, successor cor.Handler) {
		if request.Amount() < 100 {
			fmt.Printf("Amount is %v so cor1 can handle\n", request.Amount())
			return
		}
		successor.HandleRequest(request)
	}
	cor1 := cor.NewHandlerImpl(cor1HandleRequest)

	cor2HandleRequest := func(request cor.Request, successor cor.Handler) {
		fmt.Printf("Amount is %v so cor2 can handle\n", request.Amount())
	}
	cor2 := cor.NewHandlerImpl(cor2HandleRequest)

	cor1.SetSuccessor(cor2)

	cor1.HandleRequest(*req10)
	cor1.HandleRequest(*req1000)
}

func BenchmarkCorStress(b *testing.B) {

	b.StartTimer()

	req1000 := cor.NewRequest(cor.RequestType(cor.CONFERENCE), 1000)

	var prev *cor.HandlerImpl
	var first *cor.HandlerImpl
	for i := 0; i < 1000000; i++ {
		cor1HandleRequest := func(request cor.Request, successor cor.Handler) {
			if request.Amount() > 20000000 {
				fmt.Printf("Amount is %v so cor1 can handle\n", request.Amount())
				return
			}
			if successor != nil {
				successor.HandleRequest(request)
			} else {
				fmt.Printf("Can't Handle\n")
			}
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
