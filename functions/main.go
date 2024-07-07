package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func mathCalculations(x, y int) (int, int, int, int, int, string, string) {
	addition := x + y
	subtraction := x - y
	multiplication := x * y
	division := x / y
	module := x % y
	return addition, subtraction, multiplication, division, module, "Math calculations are done", "Ignored return value"
}

func main() {
	fmt.Printf("1 plus 2 is equal to: %d\n", add(1, 2))

	var function = func() {
		fmt.Println("This is an anonymous function")
	}
	function()

	var functionWithArgs = func(x string) string {
		greating := "Hello, " + x + "!"
		return greating
	}
	functionWithArgsResult := functionWithArgs("World")
	fmt.Println(functionWithArgsResult)

	addition, subtraction, multiplication, division, module, message, _ := mathCalculations(10, 5)
	fmt.Printf("Addition: %d\nSubstraction: %d\nMultiplication: %d\nDivision: %d\nModule: %d\nMessage: %s\n", addition, subtraction, multiplication, division, module, message)
}
