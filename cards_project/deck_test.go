package main

import (
	"os"
	"strings"
	"testing"
)

// Test to ensure newDeck creates a deck of 52 cards
func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}
}

// Test to make sure there are 13 cards of each suit
func TestDeckSuitCount(t *testing.T) {
	d := newDeck()
	suitCount := map[string]int{
		"Spades":   0,
		"Hearts":   0,
		"Diamonds": 0,
		"Trifoi":   0,
	}
	for _, card := range d {
		parts := strings.Split(card, " of ")
		if len(parts) != 2 {
			t.Errorf("Card format incorrect: %v", card)
			continue
		}
		suit := parts[1]
		suitCount[suit]++
	}
	for suit, count := range suitCount {
		if count != 13 {
			t.Errorf("Expected 13 cards of %v, but got %v", suit, count)
		}
	}
}

// Test to ensure that first and last card are as expected
func TestFirstAndLastCard(t *testing.T) {
	d := newDeck()

	if d[0] != "Two of Spades" {
		t.Errorf("Expected first card to be Two of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Ace of Trifoi" {
		t.Errorf("Expected last card to be Ace of Trifoi, but got %v", d[len(d)-1])
	}
}

// Test the functionality of saving to the file and making sure that there is a before and after for making sure that the file saved is deleted correctly
func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("test_deck")

	d := newDeck()
	d.saveToFile("test_deck")
	loadedDeck := newDeckFromFile("test_deck")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(loadedDeck))
	}
	if d[0] != loadedDeck[0] || d[len(d)-1] != loadedDeck[len(loadedDeck)-1] {
		t.Errorf("Expected first and last cards to match after loading from file")
	}

	os.Remove("test_deck")
}
