package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) > 52 {
		t.Errorf("Deck length should not be greater than 52")
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card as %s but got %s", "Ace of Spades", d[0])
	}
	if d[len(d)-1] != "Three of Club" {
		t.Errorf("Expected last card as %s but got %s", "Three of Club", d[len(d)-1])
	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	d := newDeck()
	d.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")
	if len(loadedDeck) > 52 {
		t.Errorf("Deck length should not be greater than 52")
	}
	os.Remove("_decktesting")

}
