package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	d := newDeck()

	if len(d) != 3 {
		t.Errorf("Expected deck length of 3, but got %v", len(d))
	}
}

func TestFirstCard(t *testing.T) {

	d := newDeck()

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}
}

func TestSaveFileToLoadFromFile(t *testing.T) {
	var filename string = "_decktesting"

	os.Remove(filename)

	d := newDeck()
	d.saveToFile(filename)

	fileDeck := loadFromFile(filename)

	if len(fileDeck) != 3 {
		t.Errorf("Expected 3 cards, but got %v", len(fileDeck))
	}

	os.Remove(filename)
}
