package main

import "fmt"

func main() {

	type Contact struct {
		pn int
	}

	type person struct {
		fn string
		ln string
		contact Contact
	}

	fmt.Println(person{ fn: "yash", ln: "soni", contact: Contact{ pn: 7385510285 },})
}
