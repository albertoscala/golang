package main

import (
	"fmt"
)

func even_or_not(n int) bool {
	if n%2 != 0 {
		return false
	}
	return true
}

func main() {
	var n int
	fmt.Println("Enter a number: ")
	fmt.Scanln(&n)
	fmt.Println(even_or_not(n))
}
