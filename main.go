package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contactInfo
}

func main() {
	jim := person {
		firstName: "Jim",
		lastName: "Ntare",
		contactInfo: contactInfo{
			email: "jim@gmail.com",
			zipCode: 94000,
		},
	}

	// jimPointer := &jim // the & is to give us the memory address of the value this variable is pointing at
	jim.updateName("Jimmy")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {// when the * is pointing to a type, it is to give us a description to a type 
	(*pointerToPerson).firstName = newFirstName // the * is to give us the value this memory addres is pointing at
}

func (p person) print() {
	fmt.Printf("%+v", p)
}