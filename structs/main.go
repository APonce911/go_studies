package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email string
	zipCode int
}

func main() {
	jim := &person{
		firstName: "Jim",
		lastName: "Joe",
		contactInfo: contactInfo{
			email: "jj@gmail.com",
			zipCode: 06671,
		},
	}

	fred := "Fred"
	fmt.Println(&fred)

	jim.print()
	jim.updateName("Bob")
	jim.print()
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
func (p person) print() {
	fmt.Println(p)
}
