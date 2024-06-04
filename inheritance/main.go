package main

import "fmt"

type person struct {
	name string
	age  uint8
}

type student struct {
	person
	course     string
	university string
}

func main() {
	person := person{name: "John Doe", age: 25}
	student := student{person, "Computer science", "MIT"}
	fmt.Println(person, student)
	fmt.Printf("Name: %s Age: %d Course: %s University: %s\n", student.person.name, student.age, student.course, student.university)
}
