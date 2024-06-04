package functions

import (
	"errors"
	"fmt"
)

func Datatypes() {
	var number int8 = 127
	var negativeNumber int = -127

	var unsignedNumber uint = 255

	// rune is an alias for int32
	var number2 rune = 127

	// byte is an alias for uint8
	var number3 byte = 255

	var floatNumber float32 = 1.234567
	var floatNumber2 float64 = 12345.6789
	floatNumber3 := 12345.6789

	var text string = "Text"
	text2 := "Text2"

	char := 'A'
	var boolean bool = true

	var emptyText string
	var zero int
	var zeroFloat float32
	var undefinedBool bool
	var errorExample error

	var definedError error = errors.New("this is an error")

	fmt.Println(number, negativeNumber, unsignedNumber, number2, number3, floatNumber, floatNumber2, floatNumber3)
	fmt.Println(text, text2)
	fmt.Println(char, boolean)
	fmt.Println(emptyText, zero, zeroFloat, undefinedBool, errorExample)
	fmt.Println(definedError)
}

// 128	64	32	16	8		4		2		1
// 1		1		1		1		1		1		1		1
