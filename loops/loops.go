package main

import (
	"fmt"
	"time"
)

func forLikeWhile() {
	i := 0

	for i < 10 {
		i++
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}

func forLikeFor() {
	for i := 1; i < 11; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}

func forWithRange() {
	persons := []string{"Alice", "Bob", "Charlie", "David", "Eve"}

	for i, person := range persons {
		fmt.Printf("Index: %d, Hello %s\n", i, person)
	}

	fmt.Println()

	for i, letter := range "Word" {
		fmt.Printf("Index: %d, Letter: %c, Code: %d\n", i, letter, letter)
	}

	fmt.Println()

	user := map[string]string{
		"age":  "25",
		"name": "Alice",
		"city": "New York",
	}

	for key, value := range user {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}

	for {
		fmt.Println("Infinite loop")
		time.Sleep(1 * time.Second)
	}
}

func main() {
	forLikeFor()
	forLikeWhile()
	forWithRange()
}
