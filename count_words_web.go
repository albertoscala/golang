package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func countWordsWeb(w http.ResponseWriter, r *http.Request) {
	counter := make(map[string]int)

	file, err := io.ReadAll(r.Body)
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

	result := ""

	for key, value := range counter {
		result += "The word " + key
		s := strconv.Itoa(value)
		result += " is in the text " + s
		result += " times!\n"
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/text")
	_, err = w.Write([]byte(result))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/countWordsWeb", countWordsWeb)

	log.Fatal(http.ListenAndServe(":8090", nil))
}
