package main

//import "fmt"

func main() {

	//cards := newDeck()
	cards := newDeckFromFile("card_deck")
	cards.shuffle()
	cards.print()
	// hand, remainingDeck := deal(cards, 5)
	// hand.print()
	// remainingDeck.print()
}
