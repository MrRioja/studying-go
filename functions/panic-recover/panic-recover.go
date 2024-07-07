package main

import "fmt"

func recoverFunction() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic:", r)
	}
}

func isStudentApproved(a, b float32) bool {
	defer recoverFunction()
	average := (a + b) / 2

	if average > 6 {
		return true
	} else if average < 6 {
		return false
	}

	panic("The average is 6. Exactly 6 is not allowed")
}

func main() {
	fmt.Println(isStudentApproved(6, 6))
	fmt.Println("End of the program")
}
