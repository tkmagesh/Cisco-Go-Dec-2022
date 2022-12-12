/* higher order functions - functions as return values */

package main

import "fmt"

func main() {
	fn := getFn(1)
	fn()

	fn = getFn(2)
	fn()

	fn = getFn(100)
	fn()

	getFn(1)()
}

func getFn(choice int) func() {
	if choice == 1 {
		return f1
	}
	if choice == 2 {
		return f2
	}
	return func() {
		fmt.Println("anon fn invoked")
	}
}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}
