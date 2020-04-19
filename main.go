package main

import (
	"fmt"

	"github.com/OlaleyeJumoke/GoInterfaceAndStruct/organization"
)

func main() {
	p := organization.NewPerson("Jumoke", "Olaleye")
	fmt.Println(p)
	fmt.Println(p.FullName())
	err := p.SetTwitterHandler("@jumoke")
	if err != nil {
		fmt.Printf("An error occur while setting twitter handler %s\n", err.Error())
	}
	println(p.TwitterHandler())
	println(p.TwitterHandler().RedirectURL())

}
