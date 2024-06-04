package main

import (
	"fmt"
	firstPackage "modulo/first-package"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Writing from main.go file in packages folder")
	firstPackage.Write()
	validation1 := checkmail.ValidateFormat("fakeemail@gmail.com")
	validation2 := checkmail.ValidateFormat("invalid-email")
	fmt.Println(validation1, validation2)
}
