package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Exam struct {
	Grade   int `json:"grade"`
	Credits int `json:"credits"`
}

func gpaCalculatorWeb(w http.ResponseWriter, r *http.Request) {
	file, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var exams []Exam

	_ = json.Unmarshal(file, &exams)

	gpa := fmt.Sprintf("%f", gpaCalc(exams))

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/text")
	_, err = w.Write([]byte(gpa + "\n"))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/gpaCalculatorWeb", gpaCalculatorWeb)

	log.Fatal(http.ListenAndServe(":8090", nil))
}
