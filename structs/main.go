package main

import "fmt"

type user struct {
	name    string
	age     uint8
	address address
}

type address struct {
	city, country string
}

func main() {
	var emptyUser user
	var user1 user = user{age: 30, name: "John Doe"}
	var user2 user = user{"Jane Doe", 25, address{}}
	user3 := user{age: 30, name: "July Doe"}
	user4 := user{"June Doe", 23, address{}}
	userWithoutAge := user{name: "January Doe"}
	userWithAddress := user{"Jenny Doe", 25, address{"Paris", "France"}}

	fmt.Println(emptyUser, user1, user2, user3, user4, userWithoutAge, userWithAddress)
}
