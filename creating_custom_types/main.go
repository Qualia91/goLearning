package main

import (
	"creating_custom_types/organization"
	"fmt"
)

func main() {
	p := organization.NewPerson("Nick", "Wood")
	println(p.ID())
	println(p.GetFirstName())
	println(p.GetLastName())
	err := p.SetTwitterHandler("@jam")
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
	println(p.GetTwitterHandler())
	println(p.GetTwitterHandler().RedirectUrl())
	fmt.Printf("Type: %T\n", p.GetTwitterHandler())
}
