package main

import (
	"io"
	"log"
	"net/http"
)

func postGreetings(w http.ResponseWriter, r *http.Request) {
	name, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.WriteString(w, "Hi "+string(name)+"!\n")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", postGreetings)

	log.Fatal(http.ListenAndServe(":8090", nil))
}
