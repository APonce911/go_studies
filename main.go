package main
import "fmt"

func main() {
  fmt.Println("Shuffle up and deal!")

  card := newCard()
  fmt.Println(card)
}

func newCard() string {
  return "Five of Diamonds"
}
