package main

import "fmt"

func genericParam(generic interface{}) {
	fmt.Println(generic)
}

func main() {
	genericParam(1)
	genericParam("Hello")
	genericParam(1.5)
	genericParam([]int{1, 2, 3})
	genericParam(map[string]int{"one": 1, "two": 2})
}
