package main

import "fmt"

func main() {
	d := NewDeck()
	fmt.Println(d.ToString())
	hand, remaining := d.Deal(5)
	hand.Print()
	remaining.Print()
}
