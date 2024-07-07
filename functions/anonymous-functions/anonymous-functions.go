package main

import "fmt"

func main() {
	func(text string) {
		println(text)
	}("Hello, Anonymous!")

	func(text string) {
		println(text)
	}("Hello, Anonymous with param!")

	returnedValue := func(text string) string {
		return fmt.Sprintf("Hello, %s!", text)
	}("Anonymous with return")

	println(returnedValue)
}
