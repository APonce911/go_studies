package main
import "fmt"

var third_card string
var forth_card string = "Queen of Clubs"

func main() {
  fmt.Println("Shuffle up and deal!")

  var card string = "Ace of Spades"
  second_card := "Ace of Diamonds"
  third_card = "Queen of Hearts"
  fmt.Println(card)
  fmt.Println(second_card)
  fmt.Println(third_card)
  fmt.Println(forth_card)
}
