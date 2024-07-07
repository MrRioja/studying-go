package main

func closure() func() {
	text := "Inside the closure"
	function := func() {
		println(text)
	}
	return function
}

func main() {
	newFunction := closure()
	newFunction()
}
