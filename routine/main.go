package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string {
		"https://searce.udemy.com/course/go-the-complete-developers-guide/learn/lecture/7797426#questions",
		"https://mail.google.com/mail/u/0/#search/mike+gouline/FMfcgzGqPphQGzdBBKjWMnThbTJZggKn",
	}

	for _,link := range links {
		go check(link)
	}
}

func check(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is down")
	}
		fmt.Println(link, "is up")
}