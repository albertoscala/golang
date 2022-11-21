package main

import (
	"io"
	"log"
	"net/http"
)

func greetings(w http.ResponseWriter, r *http.Request) {
	var name string = r.URL.Query().Get("name")
	io.WriteString(w, "Hi "+name+"!\n")
}

func main() {
	http.HandleFunc("/", greetings)

	log.Fatal(http.ListenAndServe(":8090", nil))
}
