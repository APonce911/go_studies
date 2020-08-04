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
	jim := person{
		firstName: "Jim",
		lastName: "Joe",
		contactInfo: contactInfo{
			email: "jj@gmail.com",
			zipCode: 06671,
		},
	}

	fmt.Printf("%+v", jim)
}
