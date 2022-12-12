/* higher order functions - functions as return values (applied) */

package main

import (
	"fmt"
	"log"
	"reflect"
	"runtime"
)

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

	subtract := getLogOperation(subtract)
	subtract(100, 200)

}

func getLogOperation(operationFn func(int, int)) func(int, int) {
	return func(x, y int) {
		funcName := runtime.FuncForPC(reflect.ValueOf(operationFn).Pointer()).Name()
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
