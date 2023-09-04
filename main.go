package main

import (
	"deck/deck"
	"fmt"
)

func main() {
	cardDeck := deck.New()
	deck.AddJokersToDeck(cardDeck, 2, 0)
	deck.Shuffle(cardDeck)
	deck.FilterCards(cardDeck, deck.Three)
	deck.SortDeck(cardDeck)
	for _, val := range *cardDeck {
		fmt.Printf("Suit: %s, Value: %s\n", val.Suit, val.Value)
	}
}