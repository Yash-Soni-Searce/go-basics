package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newdeck() deck {
	cards := deck{}
	cardsuits := deck{"spades", "club", "diamonds", "hearts"}
	cardvalue := deck{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "jack", "queen", "king"}
	for _, suit := range cardsuits {
		for _, value := range cardvalue {
			cards = append(cards, suit+" of "+value)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handsize int) (deck, deck) {
	cards := d[:handsize]
	remaining := d[handsize:]

	return cards, remaining
}

func (d deck) tobytestring() []byte {
	return []byte(strings.Join([]string(d), ","))
}

func (d deck) savefile(filename string, bytestring []byte) error {
	return ioutil.WriteFile(filename, bytestring, 0666)
}

func newdeckfromfile(filename string) deck {
	read, err := ioutil.ReadFile("my-cards")
	if err != nil {
		fmt.Println("problem opening the file:", err)
		os.Exit(1)
	}
	s := string(read)
	return strings.Split(s, ",")
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newposition := r.Intn(len(d) - 1)
		d[i], d[newposition] = d[newposition], d[i]
	}
}

