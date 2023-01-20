package main

import (
	"io"
	"log"
	"net/http"
)

func getGreetings(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	_, err := io.WriteString(w, "Hi "+name+"!\n")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", getGreetings)

	log.Fatal(http.ListenAndServe(":8090", nil))
}
