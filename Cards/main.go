package main

import (
	"fmt"
)

func main() {
	// var card string = "Ace of Spades"
	cards := []string{"Ace of Diamonds", newcard()}
	cards = append(cards, "Six of shades")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newcard() string {
	return "Five of Demonds"
}
