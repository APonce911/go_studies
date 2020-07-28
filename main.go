package main

import "fmt"

func main() {
	cards := newDeck()

	fmt.Println("Shuffle up and deal!")
	hand, remainingDeck := cards.deal(2)
	hand.print()
	remainingDeck.print()
}
