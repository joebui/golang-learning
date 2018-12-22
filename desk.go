package main

import "fmt"

// Desk type and functions
type desk []string

func (d desk) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}
// Desk type and functions

func newDesk() desk {
	cards := desk{}
	cardSuits := []string{"H", "D", "S", "C"}
	cardValues := []string{"ace", "1", "2", "3"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func deal(d desk, handSize int) (desk, desk) {
	return d[:handSize], d[handSize:]
}
