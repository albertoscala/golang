package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var n = 0

	for (rand.Int() % 2) != 0 {
		n++
	}

	fmt.Println("Even in ", n, " attemps")
}
