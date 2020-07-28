package main

import "fmt"

func main() {
	fmt.Println("Shuffle up and deal!")

	cards := newDeck()

	cards.print()
}
