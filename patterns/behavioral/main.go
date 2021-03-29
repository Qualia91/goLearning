package main

import "behavioral/cor"

func main() {

	// chain of responsibility
	dir := cor.NewDirector()
	vp := cor.NewVP()

	dir.SetSuccessor(vp)

	req10 := cor.NewRequest(cor.RequestType(cor.CONFERENCE), 10)
	req1000 := cor.NewRequest(cor.RequestType(cor.CONFERENCE), 1000)

	dir.HandleRequest(*req10)
	dir.HandleRequest(*req1000)
}
