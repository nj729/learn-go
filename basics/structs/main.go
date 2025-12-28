package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email       string
	phoneNumber int
}

func main() {
	// 1st method -- NOT RECOMMENDED
	// john := person{"john","wick"}
	// 2nd method
	// john := person{
	// 	firstName: "john",
	// 	lastName:  "wick",
	// }
	// 3rd method
	var john person
	john.firstName = "John"
	john.lastName = "Wick"
	john.email = "john.wick@somemail.com"
	john.phoneNumber = 1234567890
	fmt.Println(john)
	john.updateFirstName("jane")
	john.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateFirstName(name string) {
	(*p).firstName = name
}
