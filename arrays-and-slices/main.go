package main

import "fmt"

func main() {
	var array1 [5]int
	array1[0] = 1
	array1[1] = 2
	array1[2] = 3
	array1[3] = 4
	array1[4] = 5

	array2 := [5]int{5, 4, 3, 2, 1}
	array3 := [...]int{1, 3, 5, 7, 9}

	slice1 := []int{2, 4, 6, 8, 10, 12, 14}
	slice2 := array3[4:5]
	fmt.Println(array1, array2, array3, slice1, slice2)
	slice1 = append(slice1, 16)
	slice2[0] = 10
	fmt.Println(array1, array2, array3, slice1, slice2)
	fmt.Println()
	slice3 := make([]float32, 10, 15)
	fmt.Printf("Array: %v | Size: %d | Capacity: %d \n", slice3, len(slice3), cap(slice3))
}
