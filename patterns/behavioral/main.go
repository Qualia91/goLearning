package main

import (
	"behavioral/cor"
	"fmt"
)

func main() {

	// chain of responsibility
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
