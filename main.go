package main

import (
	"fmt"
	"strings"
)

// visible everywhere
var name = "alex"

func main() {

	// visible within the function
	var user string
	user = "alex"
	fmt.Print(user)

	// variable that never change use const

	const name = 123
	fmt.Print(name)
	// short variable declaration
	uni := "aleb"
	fmt.Println("Hello World", uni)

	//declare multiple variables in a single line
	age, name1 := 23, "23"
	fmt.Print(age, name1)

	// golang is a typed language
	var student string = "test"

	/*
		Basic Types
		Integer -> int, int8, int16, int32, int64, rune,uint,uintptr,uint8,uint16,uint32, uint64
		unit - unsigned int - double amt of values if no is not negative
		Float -> float32, float64
		Complex -> complex64, complex128
		Byte -> byte - single ASCII characters
		Strings -> sequence of bytes - are immutable
		Booleans -> bool /true/false


	*/
	fmt.Println(len(student))
	fmt.Println(student[0])
	fmt.Println(student[0:2])
	strings.HasPrefix("st", "te")

	/*
		strings are reference types
		-if passing a string to a function, reference to the string will be copied not its value
	*/

	/*
		Arrays -> sequence of items of a single type
		Are value types - copying an array

	*/
	var arr []string = {"alex", "m"}


	myarr := []string{"jan", "feb", "march", "july"}
	fmt.Println(myarr)
}
