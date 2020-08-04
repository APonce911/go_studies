package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	var airton person
	airton.firstName = "Airton"
	airton.lastName = "Ponce"

	fmt.Println(airton)
	fmt.Printf("%+v", airton)
}
