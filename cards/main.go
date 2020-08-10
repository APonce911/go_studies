package main

func main() {
	cards := newDeckFromFile("my_deck")
	cards.shuffle()
	cards.print()
}
