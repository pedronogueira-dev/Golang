package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

//Deck is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return append(cards, "Joker")
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, num int) (deck, deck) {
	return d[:num], d[num:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fn string) error {
	return ioutil.WriteFile(fn, []byte(d.toString()), 0666)
}

func newDeckFromFile(fn string) (deck, error) {
	fc, err := ioutil.ReadFile(fn)

	if err != nil {
		return nil, fmt.Errorf("File Read Error: %s", err)
	}

	cstring := string(fc)
	return deck(strings.Split(cstring, ",")), nil
}

func (d deck) shuffle() {
	src := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(src)
	for i := range d {
		pos := rnd.Intn(len(d) - 1)
		d[i], d[pos] = d[pos], d[i]
	}

}
