package main

import "fmt"

func firstFunction() {
	fmt.Println("First function")
}

func secondFunction() {
	fmt.Println("Second function")
}

func isStudentApproved(a, b float32) bool {
	defer fmt.Println("Average calculated. Result will be returned")
	fmt.Println("Checking if student is approved")

	average := (a + b) / 2

	if average >= 6 {
		return true
	}

	return false
}

func main() {
	firstFunction()
	defer secondFunction()
	fmt.Println(isStudentApproved(6, 5))
}
