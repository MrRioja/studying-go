package main

import "fmt"

func main() {
	user := map[string]string{
		"username": "johndoe",
		"name":     "John Doe",
		"email":    "johndoe@email.com",
	}

	fmt.Println(user)
	fmt.Printf("Name: %s | Username: %s | Email: %s\n", user["name"], user["username"], user["email"])

	users := map[string]map[string]string{
		"johndoe": user,
	}

	fmt.Println(users)
	fmt.Printf("Value of key 'johndoe': %v\n", users["johndoe"])
	delete(users, "johndoe")
	fmt.Printf("Map after remove key 'johndoe': %v\n", users)
}
