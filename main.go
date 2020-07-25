package main

import "fmt"

func main() {
	fmt.Println("Shuffle up and deal!")

	cards := []string{newCard(), "Ace of Diamonds"}
	cards = append(cards, "Six of Spades")
	fmt.Println(cards)
	var cardsArray [3]int = [3]int{3, 2, 1}
	fmt.Println(cardsArray)

	for index, card := range cardsArray {
		fmt.Println(index, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
