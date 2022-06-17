package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Error("Expected deck length of 20, but got ", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Error("Expected first card of Ace of Spades, but got ", d[0])
	}

}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Error("Expected 16 cards in deck, got ", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
