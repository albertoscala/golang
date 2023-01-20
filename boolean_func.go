package main

import (
	"fmt"
)

func evenOrNot(n int) bool {
	if n%2 != 0 {
		return false
	}
	return true
}

func main() {
	var n int
	fmt.Println("Enter a number: ")
	fmt.Scanln(&n)
	fmt.Println(evenOrNot(n))
}
