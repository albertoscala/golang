package main

import (
	"fmt"
	"time"
)

func main() {

	if time.Now().Second()%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}
