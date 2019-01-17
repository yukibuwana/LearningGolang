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
	yuki := person{
		firstName: "Yuki",
		lastName:  "Buwana",
		contactInfo: contactInfo{
			email:   "yukibuwana@gmail.com",
			zipCode: 60237,
		},
	}

	fmt.Println(&yuki)
	yuki.updateName("Ayunda Maulida")
	yuki.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
