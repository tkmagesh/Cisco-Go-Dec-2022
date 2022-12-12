/* function type declaration*/

package main

import (
	"fmt"
	"log"
	"reflect"
	"runtime"
	"time"
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

	//composing "log" operation
	/*
		add := getLogOperation(add)
		add(100, 200)

		subtract := getLogOperation(subtract)
		subtract(100, 200)
	*/

	//composing "profile" operation
	/*
		profileAdd := getProfileOperation(add)
		profileAdd(100, 200)

		profileSubtract := getProfileOperation(subtract)
		profileSubtract(100, 200)
	*/

	//composing both "log" & "profile" operations
	logAdd := getLogOperation(add)
	profileLogAdd := getProfileOperation(logAdd)
	profileLogAdd(100, 200)

	getProfileOperation(getLogOperation(subtract))(100, 200)
}

func getProfileOperation(operationFn Operation) Operation {
	funcName := runtime.FuncForPC(reflect.ValueOf(operationFn).Pointer()).Name()
	return func(i1, i2 int) {
		start := time.Now()
		operationFn(i1, i2)
		elapsed := time.Since(start)
		fmt.Printf("Time taken for [%q] = %v\n", funcName, elapsed)
	}
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
