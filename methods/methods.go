package main

import "fmt"

type user struct {
	name     string
	email    string
	notified bool
}

func (u *user) notify() {
	println("Sending user email to", u.name, "<", u.email, ">")
	u.notified = true
}

func (u *user) changeEmail(email string) {
	u.email = email
}

func main() {
	u := user{"John Doe", "johndoe@email.com", false}
	fmt.Println(u)
	u.changeEmail("johnjoe@changed.com")
	u.notify()
	fmt.Println(u)
}
