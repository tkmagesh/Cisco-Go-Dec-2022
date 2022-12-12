/* higher order functions - functions as arguments (applied) */

package main

import (
	"fmt"
	"log"
)

func main() {
	/*
		log.Println("operation started")
		add(100, 200)
		log.Println("operation completed")

		log.Println("operation started")
		subtract(100, 200)
		log.Println("operation completed")
	*/

	/*
		logAdd(100, 200)
		logSubtract(100, 200)
	*/

	logOperation(100, 200, add)
	logOperation(100, 200, subtract)
	logOperation(100, 200, func(x, y int) {
		fmt.Println("Multiply Result :", x*y)
	})

}

func logOperation(x, y int, operationFn func(int, int)) {
	log.Println("operation started")
	operationFn(x, y)
	log.Println("operation completed")
}

func logAdd(x, y int) {
	log.Println("operation started")
	add(x, y)
	log.Println("operation completed")
}
func logSubtract(x, y int) {
	log.Println("operation started")
	subtract(x, y)
	log.Println("operation completed")
}

//3rd party api
func add(x, y int) {
	fmt.Println("Add Result :", x+y)
}
func subtract(x, y int) {
	fmt.Println("Subtract Result :", x-y)
}
