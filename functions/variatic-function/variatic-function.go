package main

func sumNumbers(numbers ...int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func main() {
	println(sumNumbers(1, 2, 3, 4, 5))
	println(sumNumbers(1, 2, 3))
	println(sumNumbers(1, 2))
	println(sumNumbers(1))
	println(sumNumbers())
}
