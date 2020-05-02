package main

import (
	"fmt"
	"strings"
)

type Deck []string

func NewDeck() Deck {
	d := Deck {}
	suits := []string { "Spades", "Diamonds", "Hearts", "Clubs" }
	values := []string { "Ace", "Two", "Three", "Four" }
	for _, suit := range suits {
		for _, value := range values {
			d = append(d, value + " of " + suit)
		}
	}
	return d
}

func (d Deck) Print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d Deck) Deal(handSize int) (Deck, Deck) {
	return d[:handSize], d[handSize:]
}

func (d Deck) ToString() string {
	return strings.Join(d, ",")
}

func (d Deck) SaveToFile(filename string) error {

}