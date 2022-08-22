package main

import "fmt"

type bot interface {
	getgreeting() string
}

type englishbot struct {}

type spanishbot struct {}

func printing(a bot) {
	fmt.Println(a.getgreeting())
}

func (englishbot) getgreeting() string {
	return "hii"
}

func (spanishbot) getgreeting() string {
	return "hola"
}

func main() {
	eb := englishbot{}
	sb := spanishbot{}

	printing(eb)
	printing(sb)
}