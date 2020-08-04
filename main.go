package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	airton := person{firstName: "Airton", lastName: "Ponce"}
	fmt.Println(airton)
	fmt.Println(airton.firstName)
	fmt.Println(airton.lastName)
}
