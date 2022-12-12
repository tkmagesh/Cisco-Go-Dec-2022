package main

import "fmt"

/* var x int */
/* var x int = 100 */
/* var x = 100 */

/* x := 100 */

func main() {
	/*
		var msg string
		msg = "Hello World!"
	*/
	/*
		var msg string = "Hello World!"
	*/

	// type inference
	/*
		var msg = "Hello World!"
	*/

	msg := "Hello World!" // ":=" syntax cannot be used for the package scoped variables
	fmt.Println(msg)

	//Using more than one variable
	/*
		var x int
		var y int
		var str string
		var result int
		x = 100
		y = 200
		str = "Sum of 100 and 200 :"
		result = x + y
		fmt.Println(str, result)
	*/

	/*
		var x, y, result int
		var str string
		x = 100
		y = 200
		str = "Sum of 100 and 200 :"
		result = x + y
		fmt.Println(str, result)
	*/

	/*
		var (
			x, y, result int
			str          string
		)
		x = 100
		y = 200
		str = "Sum of 100 and 200 :"
		result = x + y
		fmt.Println(str, result)
	*/

	/*
		var x int = 100
		var y int = 200
		var str string = "Sum of 100 and 200 :"
		var result int = x + y
		fmt.Println(str, result)
	*/

	/*
		var x, y int = 100, 200
		var str string = "Sum of 100 and 200 :"
		var result int = x + y
		fmt.Println(str, result)
	*/

	/*
		var (
			x, y   int    = 100, 200
			str    string = "Sum of 100 and 200 :"
			result int    = x + y
		)
		fmt.Println(str, result)
	*/

	/*
		var (
			x, y   = 100, 200
			str    = "Sum of 100 and 200 :"
			result = x + y
		)
		fmt.Println(str, result)
	*/

	/*
		var (
			x, y, str = 100, 200, "Sum of 100 and 200 :"
			result    = x + y
		)
		fmt.Println(str, result)
	*/

	/*
		x, y, str := 100, 200, "Sum of 100 and 200 :"
		result := x + y
		fmt.Println(str, result)
	*/

	/*
		x, y, str := 100, 200, "Sum of 100 and 200 :"
		result := x + y
		fmt.Println(str, result)
	*/

	//using the format verbs with "Printf" function
	x, y, str := 40, 50, "Sum of %d and %d : %d\n"
	result := x + y
	fmt.Printf(str, x, y, result)
}
