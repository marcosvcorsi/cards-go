package main

import "fmt"

func main() {
	// var card string = "Ace of spades"
	//card := "Ace of spades"
	// card = "Five of Diamonds"

	// card := newCard()

	// fmt.Println(card)

	cards := []string{newCard(), "Ace of spades"}
	cards = append(cards, "Six of Diamonds")

	for i, card := range cards {
		fmt.Println(i, card)
	}

	fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamonds"
}
