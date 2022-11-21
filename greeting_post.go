package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func greetings(w http.ResponseWriter, r *http.Request) {
	name, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	io.WriteString(w, "Hi "+string(name)+"!\n")
}

func main() {
	http.HandleFunc("/", greetings)

	log.Fatal(http.ListenAndServe(":8090", nil))
}
