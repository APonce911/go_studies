package main

import "fmt"

type deck []string

func (d deck) print() {
	fmt.Println(d)

	for _, card := range d {
		fmt.Println(card)
	}
}
