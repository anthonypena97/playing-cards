package main

import "fmt"

func main() {

	// var card string = "Ace of Spades"
	// card := "Ace of Spades" // variable with type String
	// card = "Five of Diamonds"

	// card := newCard()

	// fmt.Println(card)

	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	// iterating over a closed set
	for i, card := range cards {
		fmt.Println(i, card)
	}

	// fmt.Println(cards)

}

func newCard() string {

	return "Five of Diamonds"

}
