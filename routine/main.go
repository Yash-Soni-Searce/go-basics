package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string {
		"https://searce.udemy.com/course/go-the-complete-developers-guide/learn/lecture/7797426#questions",
		"https://mail.google.com/mail/u/0/#search/mike+gouline/FMfcgzGqPphQGzdBBKjWMnThbTJZggKn",
		"https://notog.org",
	}

	c := make(chan string)

	for _,link := range links {
		go check(link, c)
		fmt.Println(<- c)
	}
	close(c)
}

func check(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is down")
	} else {
		fmt.Println(link, "is up")
		c <- "Yes it is up"
	}
}