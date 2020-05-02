package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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

func NewDeckFromFile(filename string) Deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return Deck(s)
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
	return ioutil.WriteFile(filename, []byte(d.ToString()), 0666)
}

func (d Deck) Shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}