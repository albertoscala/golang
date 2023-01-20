package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func readGrade() string {
	var result string

	fmt.Printf("Insert grade/credits: ")
	_, err := fmt.Scanf("%s", &result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func gpaCalc(exams []Exam) float32 {
	gpa := float32(0)
	cfu := float32(0)

	for _, exam := range exams {
		gpa += float32(exam.Grade * exam.Credits)
		cfu += float32(exam.Credits)
	}

	return gpa / cfu
}

func main() {
	var result string

	gpa := float32(0)
	credits := 0

	result = readGrade()
	for result != "-1" {
		arr := strings.Split(result, "/")
		grade, _ := strconv.Atoi(arr[0])
		credit, _ := strconv.Atoi(arr[1])
		gpa += float32(grade * credit)
		credits += credit
		result = readGrade()
	}

	fmt.Println("Your GPA is: ", gpa/float32(credits))
}
