package main

func invertSign(x int) int {
	return -x
}

func invertSignWithPointer(x *int) {
	*x = -*x
}

func main() {
	number := 20
	invertedNumber := invertSign(number)
	println(invertedNumber, number)
	newNumber := 40
	invertSignWithPointer(&newNumber)
	println(newNumber)
}
