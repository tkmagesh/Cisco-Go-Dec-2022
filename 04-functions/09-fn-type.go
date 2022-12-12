/* function type declaration*/

package main

import (
	"fmt"
	"log"
	"reflect"
	"runtime"
)

type Operation func(int, int)

func main() {
	/*
		add(100, 200)
		subtract(100, 200)
	*/
	/*
		logOperation(100, 200, add)
		logOperation(100, 200, subtract)
	*/

	add := getLogOperation(add)
	add(100, 200)
	add(10, 20)
	add(1000, 2000)
	/* subtract := getLogOperation(subtract)
	subtract(100, 200) */

}

func getLogOperation(operationFn Operation) Operation {
	fmt.Println("Deriving the function name")
	funcName := runtime.FuncForPC(reflect.ValueOf(operationFn).Pointer()).Name()
	return func(x, y int) {
		log.Printf("[%q] operation started\n", funcName)
		operationFn(x, y)
		log.Printf("[%q] operation completed\n", funcName)
	}
}

//3rd party api
func add(x, y int) {
	fmt.Println("Add Result :", x+y)
}
func subtract(x, y int) {
	fmt.Println("Subtract Result :", x-y)
}
