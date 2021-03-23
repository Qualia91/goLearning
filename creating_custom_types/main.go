package main

import (
	"creating_custom_types/organization"
	"fmt"
)

func main() {
	p := organization.NewPerson("Nick", "Wood", organization.NewEuropeanUnionIdNumber("UK", "123-456-789"))
	p2 := organization.NewPerson("Nick", "Wood", organization.NewEuropeanUnionIdNumber("UK", 123))
	println(p.ID())
	println(p.GetFirstName())
	println(p.GetLastName())
	println(p2.ID())
	println(p2.GetFirstName())
	println(p2.GetLastName())

	err := p.SetTwitterHandler("@jam")
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
	println(p.GetTwitterHandler())
	println(p.GetTwitterHandler().RedirectUrl())
	fmt.Printf("Type: %T\n", p.GetTwitterHandler())
	fmt.Println(p.Citizen.Country())

	name1 := NameTest{
		firstName: "Nick",
		lastName:  "Wood",
	}
	name2 := NameTest{
		firstName: "Nick",
		lastName:  "Wood",
	}

	if name1 == name2 {
		println("We match because NameTest has a simple memory layout and so go can compare the values in the object")
	}

}

type NameTest struct {
	firstName string
	lastName  string
}
