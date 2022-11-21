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

func gpa_calc(exams []Exam) float32 {
	gpa := float32(0)
	cfu := float32(0)

	for _, exam := range exams {
		gpa += float32(exam.Grade * exam.Credits)
		cfu += float32(exam.Credits)
	}

	return gpa / cfu
}

func gpaCalculatorWeb(w http.ResponseWriter, r *http.Request) {
	file, _ := io.ReadAll(r.Body)

	exams := []Exam{}

	_ = json.Unmarshal(file, &exams)

	gpa := fmt.Sprintf("%f", gpa_calc(exams))

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/text")
	w.Write([]byte(gpa + "\n"))

	return
}

func main() {
	http.HandleFunc("/gpaCalculatorWeb", gpaCalculatorWeb)

	log.Fatal(http.ListenAndServe(":8090", nil))
}
