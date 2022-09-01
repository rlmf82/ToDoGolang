package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
	contact   contactInfo
}

type contactInfo struct {
	email   string
	zipCode string
}

func createPerson() {

	var alice person
	alice.firstName = "Alice"
	alice.lastName = "Farias"
	alice.age = 1

	aliceContact := contactInfo{
		email:   "alice@farias.com",
		zipCode: "2700130",
	}

	alice.contact = aliceContact
	alice.print()

	//pointer := &alice
	alice.updateName("Lili")
	alice.print()
}

//This function receives the pointer to person.
func (p *person) updateName(newName string) {
	p.firstName = newName
}

//This function receives the person value, not reference(pointer).
func (p person) print() {
	fmt.Printf("%+v", p)
	fmt.Println()
}
