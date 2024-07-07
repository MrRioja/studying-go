package main

func fibonacci(n uint) uint {
	if n < 2 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	println(fibonacci(1000))
}
