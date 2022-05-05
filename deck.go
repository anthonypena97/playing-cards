package main

import (
	"fmt"
	"strings"
)

// Create a new type of 'deck'
// which is a slice of string
type deck []string

func newDeck() deck {

	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// 'i' and 'j' are replaced with underscores, '_' , to let go know we don't need these variables
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards

}

// function with a receiver declared as 'd'
// by convention - the receiver is commonly decalred as one or two letters that match to the type
func (d deck) print() {

	for i, card := range d {
		fmt.Println(i, card)
	}

}

// the second paranthesis are the two types we will be retrning
func deal(d deck, handSize int) (deck, deck) {

	return d[:handSize], d[handSize:]

}

func (d deck) toString() string {

	return strings.Join([]string(d), ", ")

}
