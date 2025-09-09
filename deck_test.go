package main

import (
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
