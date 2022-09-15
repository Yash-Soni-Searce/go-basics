package main

import "fmt"

type bot interface {
	getgreeting(x string, y string) string
}

type GCP struct{
	a string
	b string
}

type AWS struct{
	a string
}

func printing(a bot) {
	fmt.Println(a.getgreeting())
}

func (GCP) getgreeting() string {
	return "hii"
}

func (AWS) getgreeting() string {
	return "hola"
}

func main() {
	eb := GCP{}
	sb := AWS{}

	printing(eb)
	printing(sb)
}
//check for updated name
