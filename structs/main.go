package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contactInfo: contactInfo{
			email:   "alex@gmail.com",
			zipCode: 123456,
		},
	}
	alex.print()
	alex.updateName("Madhav")
	alex.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateName(newName string) {
	(*p).firstName = newName
}
