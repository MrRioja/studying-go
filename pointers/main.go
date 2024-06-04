package main

import "fmt"

func main() {
	var variable int = 10
	var variable2 int = variable
	var pointer *int = &variable

	fmt.Println(variable, variable2, *pointer)
	variable++
	fmt.Println(variable, variable2, *pointer)
}
