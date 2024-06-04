package main

import "fmt"

func main() {
	number := 10

	if number > 5 {
		fmt.Println("Number is greater than 5")
	} else if number < 5 {
		fmt.Println("Number is less than or equal to 5")
	} else {
		fmt.Println("Number is equal to 5")
	}

	if otherNumber := number; otherNumber%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}
}
