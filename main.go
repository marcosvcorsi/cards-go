package main

func main() {
	// var card string = "Ace of spades"
	//card := "Ace of spades"
	// card = "Five of Diamonds"

	// card := newCard()

	// fmt.Println(card)

	// cards := newDeck()

	// hand, remainingCards := deal(cards, 5)

	// hand.print()

	// remainingCards.print()

	cards := newDeck()
	cards.saveToFile("my_cards")

	newCards := newDeckFromFile("my_cards")
	newCards.print()
}
