package main

import (
	"fmt"

	"github.com/OlaleyeJumoke/GoInterfaceAndStruct/organization"
)

func main() {
	p := organization.NewPerson("Jumoke", "Olaleye", organization.NewEuropeanUnionIdentifier("123456789", "Germany"))
	fmt.Println(p)
	fmt.Println(p.FullName())
	err := p.SetTwitterHandler("@jumoke")
	if err != nil {
		fmt.Printf("An error occur while setting twitter handler %s\n", err.Error())
	}
	name1 := Name{first: "Jumoke", last: "Olaleye"}
	name2 := Name{first: "Jumoke", last: "Olaleye"}

	if name1 == name2 {
		println("we match")
	}
	
 
}

// Name strct
type Name struct {
	first string
	last  string
}
