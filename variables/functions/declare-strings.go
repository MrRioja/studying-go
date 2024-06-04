package functions

import "fmt"

func DeclareStrings() {
	var name string = "John Doe"
	name2 := "Jane Doe"

	var (
		name3 string = "Juan Doe"
		name4 string = "Jonny Doe"
	)

	name5, name6 := "Jude Doe", "July Doe"

	name7 := "Jennifer Doe"
	name7 = "Jenny Doe"

	const constantName string = "Constant Name"

	fmt.Println(name)
	fmt.Println(name2)
	fmt.Println(name3, name4)
	fmt.Println(name5, name6)

	name5, name6 = name6, name5

	fmt.Println("Inverted values: ", name5, name6)
	fmt.Println(name7)
	fmt.Println(constantName)
}
