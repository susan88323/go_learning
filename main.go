package main

func main() {
	// var card string = "Ace of Spades"
	// card := newCard()
	// fmt.Println(card)
	// cards := newDeck()

	// hand, remainingDeck := deal(cards, 5)
	// hand.print()
	// remainingDeck.print()

	cards := newDeck()
	cards.shuffle()
	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")
	// cards := newDeckFromFile("my_cards")
	cards.print()

}
