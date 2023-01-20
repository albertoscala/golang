package main

import (
	"fmt"
	"strconv"
	"strings"
)

func read_grade() string {
	var result string
	fmt.Printf("Insert grade/credits: ")
	fmt.Scanf("%s", &result)
	return result
}

func main() {
	var result string

	gpa := float32(0)
	credits := 0

	result = read_grade()
	for result != "-1" {
		arr := strings.Split(result, "/")
		grade, _ := strconv.Atoi(arr[0])
		credit, _ := strconv.Atoi(arr[1])
		gpa += float32(grade * credit)
		credits += credit
		result = read_grade()
	}

	fmt.Println("Your GPA is: ", gpa/float32(credits))
}
