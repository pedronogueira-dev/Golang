package main

import "fmt"

func main() {
	cards, err := newDeckFromFile("deck_state.txt")

	if err != nil {
		fmt.Println(err)
		cards = newDeck()
	}
	cards.print()
	//cards.saveToFile("deck_state.txt")
}
