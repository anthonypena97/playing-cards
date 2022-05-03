package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of string
type deck []string

// function with a receiver declared as 'd'
// by convention - the receiver is commonly decalred as one or two letters that match to the type
func (d deck) print() {

	for i, card := range d {
		fmt.Println(i, card)
	}

}
