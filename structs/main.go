package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
	contactInfo
}

type contactInfo struct {
	email       string
	phoneNumber string
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateFirstName(fn string) {
	p.firstName = fn
}

func main() {

	p := person{
		firstName: "Alex",
		lastName:  "Anderson",
		age:       30 ,
	}
	p.contactInfo = contactInfo{
		email:       "example.mail@ex.com",
		phoneNumber: "999 999 999",
	}

	fmt.Println(p)
	p.updateFirstName("Jim")
	p.print()
}
