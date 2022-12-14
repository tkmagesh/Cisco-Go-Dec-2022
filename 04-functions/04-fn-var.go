package main

import "fmt"

func main() {
	var sayHi func()
	sayHi = func() {
		fmt.Println("Hi there!")
	}
	sayHi()

	var greet func(string)
	greet = func(userName string) {
		fmt.Printf("Hi %q, Have a nice day!\n", userName)
	}
	greet("Magesh")

	/* declare the variables for each of the below functions and use them (as above) */
	var greetFn func(string) string
	greetFn = func(userName string) string {
		var msg string
		msg = fmt.Sprintf("Hi %q, Have a nice day", userName)
		return msg
	}
	greetMsg := greetFn("Magesh")
	fmt.Println(greetMsg)

	var add func(int, int) int
	add = func(x, y int) int {
		return x + y
	}
	result := add(100, 200)
	fmt.Println(result)

	var divide func(int, int) (int, int)
	divide = func(x, y int) (quotient, remainder int) {
		quotient = x / y
		remainder = x % y
		return
	}
	q, r := divide(100, 7)
	fmt.Printf("Dividing 100 by 7, quotient = %d and remainder = %d\n", q, r)
}
