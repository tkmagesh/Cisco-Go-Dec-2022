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
	greetMsg := func(userName string) string {
		var msg string
		msg = fmt.Sprintf("Hi %q, Have a nice day", userName)
		return msg
	}("Magesh")
	fmt.Println(greetMsg)

	result := func(x, y int) int {
		return x + y
	}(100, 200)
	fmt.Println(result)

	q, r := func(x, y int) (quotient, remainder int) {
		quotient = x / y
		remainder = x % y
		return
	}(100, 7)
	fmt.Printf("Dividing 100 by 7, quotient = %d and remainder = %d\n", q, r)
}
