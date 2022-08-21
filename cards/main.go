package main

import (
	"fmt"
)

func main() {
	// cards := newdeck()
	// cards.print()
	// fmt.Println("New deck!!!!")
	// deal, remainingdeal := deal(cards, 5)
	// deal.print()
	// remainingdeal.print()
	// fmt.Println("New new new")
	// fmt.Println(cards.tobytestring())
	// cards.savefile("my-cards", cards.tobytestring())
	// cards.newdeckfromfile("my-cards")
	cards:=newdeckfromfile("my-cards")
	cards.shuffle()
	fmt.Println("got a new deck")
	cards.print()
}
