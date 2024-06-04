package main

import "fmt"

func main() {
	fmt.Printf("Result of 10 + 5: %d\n", 10+5)
	fmt.Printf("Result of 10 - 5: %d\n", 10-5)
	fmt.Printf("Result of 10 * 5: %d\n", 10*5)
	fmt.Printf("Result of 10 / 5: %d\n", 10/5)
	fmt.Printf("Result of 10 %% 5: %d\n", 10%5)
	fmt.Println()
	fmt.Printf("Result of 1 > 2: %t\n", 1 > 2)
	fmt.Printf("Result of 1 < 2: %t\n", 1 < 2)
	fmt.Printf("Result of 1 >= 2: %t\n", 1 >= 2)
	fmt.Printf("Result of 1 <= 2: %t\n", 1 <= 2)
	fmt.Printf("Result of 1 == 2: %t\n", 1 == 2)
	fmt.Printf("Result of 1 != 2: %t\n", 1 != 2)
	fmt.Println()
	trueValue, falseValue := true, false
	fmt.Printf("Result of true && false: %t\n", trueValue && falseValue)
	fmt.Printf("Result of true || false: %t\n", trueValue || falseValue)
	fmt.Printf("Result of !true and !false: %t %t\n", !trueValue, !falseValue)
	fmt.Println()
	value := 10
	fmt.Printf("Value: %d\n", value)
	value++
	fmt.Printf("Result of value++: %d\n", value)
	value += 4
	fmt.Printf("Result of value += 4: %d\n", value)
	value--
	fmt.Printf("Result of value--: %d\n", value)
	value -= 4
	fmt.Printf("Result of value -= 4: %d\n", value)
	value *= 5
	fmt.Printf("Result of value *= 5: %d\n", value)
	value /= 5
	fmt.Printf("Result of value /= 5: %d\n", value)
	value %= 5
	fmt.Printf("Result of value %%= 5: %d\n", value)
}
