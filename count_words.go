package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	counter := make(map[string]int)

	file, err := os.ReadFile("TheArabianNights.txt")
	if err != nil {
		log.Fatal(err)
	}

	txt := string(file[:])

	// Replaces
	txt = strings.ReplaceAll(txt, "\n", " ")
	txt = strings.ReplaceAll(txt, "\"", "")
	txt = strings.ReplaceAll(txt, ",", "")
	txt = strings.ReplaceAll(txt, ".", "")
	txt = strings.ReplaceAll(txt, "?", "")
	txt = strings.ReplaceAll(txt, "!", "")
	txt = strings.ReplaceAll(txt, ";", "")
	txt = strings.ReplaceAll(txt, ":", "")

	txt = strings.ToLower(txt)

	words := strings.Split(txt, " ")

	for i := 0; i < len(words); i++ {
		if _, ok := counter[words[i]]; ok {
			counter[words[i]] += 1
		} else {
			counter[words[i]] = 1
		}
	}

	for key, value := range counter {
		fmt.Print("The word ", key)
		fmt.Print(" is in the text ", value)
		fmt.Println(" times!")
	}
}
