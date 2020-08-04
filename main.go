package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact contactInfo
}

type contactInfo struct {
	email string
	zipCode int
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName: "Joe",
		contact: contactInfo{
			email: "jj@gmail.com",
			zipCode: 06671,
		},
	}

	fmt.Println(jim)
	fmt.Printf("%+v", jim)
}
