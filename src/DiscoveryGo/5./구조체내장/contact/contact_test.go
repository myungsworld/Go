package main

import "fmt"

type Address struct {
	City  string
	State string
}

type Telephone struct {
	Mobile string
	Direct string
}

type Contact struct {
	Address
	Telephone
}

func ExampleContact() {
	var c Contact
	c.Mobile = "010-2046-7089"
	fmt.Println(c.Telephone.Mobile)
	c.Address.City = "San Francisco"
	c.State = "CA"
	fmt.Println(c)
	// Output:
	// .

}
