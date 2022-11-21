package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func getWasa(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())

	var n = 0

	for (rand.Int() % 2) != 0 {
		n++
	}

	s := fmt.Sprintf("Even in %d attemps \n", n)

	fmt.Println(s)
	io.WriteString(w, s)
}

func main() {
	http.HandleFunc("/wasa", getWasa)

	log.Fatal(http.ListenAndServe(":8090", nil))
}
