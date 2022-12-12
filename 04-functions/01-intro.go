/* function basics */

package main

import "fmt"

func main() {
	say_hi()
	greet("Magesh")
	fmt.Println(getGreetMsg("Raj"))
	fmt.Println(add(100, 200))
	// fmt.Println(divide(100, 7))

	/*
		q, r := divide(100, 7)
		fmt.Printf("Dividing 100 by 7, quotient = %d and remainder = %d\n", q, r)
	*/

	//print ONLY the quotient
	q, _ := divide(100, 7)
	fmt.Printf("Dividing 100 by 7, quotient = %d\n", q)
}

/* function with no arguments & no return results */
func say_hi() {
	fmt.Println("Hi there!")
}

/* function with 1 argument */
func greet(userName string) {
	fmt.Printf("Hi %q, Have a nice day!\n", userName)
}

/* function with 1 argument & 1 return result */
func getGreetMsg(userName string) string {
	var msg string
	msg = fmt.Sprintf("Hi %q, Have a nice day", userName)
	return msg
}

/* function with > 1 argument & 1 return result */
/*
func add(x int, y int) int {
	return x + y
}
*/

func add(x, y int) int {
	return x + y
}

/* function with > 1 argument & > 1 return result */
/*
func divide(x, y int) (int, int) {
	var quotient, remainder int
	quotient = x / y
	remainder = x % y
	return quotient, remainder
}
*/

// named return results
/*
func divide(x, y int) (quotient int, remainder int) {
	quotient = x / y
	remainder = x % y
	return
}
*/

func divide(x, y int) (quotient, remainder int) {
	quotient = x / y
	remainder = x % y
	return
}
