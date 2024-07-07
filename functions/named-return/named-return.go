package main

import "fmt"

func mathCalculations(x, y int) (addition int, subtraction int, multiplication int, division int, module int, finishMessage string, ignoredMessage string) {
	module = x % y
	addition = x + y
	division = x / y
	subtraction = x - y
	multiplication = x * y
	ignoredMessage = "Ignored return value"
	finishMessage = "Math calculations are done"

	return
}

func main() {
	fmt.Println(mathCalculations(10, 5))
}
