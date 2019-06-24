package main

import (
	"os"
	"testing"
)

//New Deck
// Has 52 Cards (exSize)

func TestNewDeck(t *testing.T) {
	exDeck := []string{"Ace of Spades", "Two of Spades", "Three of Spades", "Four of Spades", "Five of Spades", "Six of Spades",
		"Seven of Spades", "Eight of Spades", "Nine of Spades", "Ten of Spades", "Jack of Spades", "Queen of Spades", "King of Spades",
		"Ace of Hearts", "Two of Hearts", "Three of Hearts", "Four of Hearts", "Five of Hearts", "Six of Hearts",
		"Seven of Hearts", "Eight of Hearts", "Nine of Hearts", "Ten of Hearts", "Jack of Hearts", "Queen of Hearts", "King of Hearts",
		"Ace of Diamonds", "Two of Diamonds", "Three of Diamonds", "Four of Diamonds", "Five of Diamonds", "Six of Diamonds",
		"Seven of Diamonds", "Eight of Diamonds", "Nine of Diamonds", "Ten of Diamonds", "Jack of Diamonds", "Queen of Diamonds", "King of Diamonds",
		"Ace of Clubs", "Two of Clubs", "Three of Clubs", "Four of Clubs", "Five of Clubs", "Six of Clubs",
		"Seven of Clubs", "Eight of Clubs", "Nine of Clubs", "Ten of Clubs", "Jack of Clubs", "Queen of Clubs", "King of Clubs",
		"Joker"}
	d := newDeck()
	i := 0

	if len(d) != len(exDeck) {
		t.Errorf("Expected deck of length %v, but got %v", len(exDeck), len(d))
	}

	for i < len(exDeck) {
		if d[i] != exDeck[i] {
			t.Errorf("Expected position %v to be %v, but instead was %v", i+1, exDeck[i], d[i])
		}
		i++
	}

}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()

	d.saveToFile("_decktesting")

	ld, err := newDeckFromFile("_decktesting")

	if err != nil {
		t.Errorf("Error loading deck from file")
	}
	if len(ld) != len(d) {
		t.Errorf("Error saving/loading deck from file")
	}
}
