package main

import "fmt"

func main() {
	num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, count := range num {
		if count == 0 {
			fmt.Println("0 is neither odd nor even")
		} else if count%2 == 0 {
			fmt.Println(count, " is even")
		} else {
			fmt.Println(count, "is odd")
		}
	}
}
