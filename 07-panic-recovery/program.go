package main

import (
	"errors"
	"fmt"
	"runtime/debug"
)

var DivideByZeroError error = errors.New("divide by zero error")

func main() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("something went wrong!")
			fmt.Println(e)
			debug.PrintStack()
			return
		}
		fmt.Println("Thank you!")
	}()

	result, err := divideClient(100, 0)
	if err != nil {
		fmt.Println("Divide resulted in an error. Take appropriate action")
	} else {
		fmt.Println("Result :", result)
	}
}

func divideClient(x, y int) (result int, err error) {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("[deferred divideClient], panic data :", e)
			/*
				if er, ok := e.(error); ok {
					err = er
					return
				}
			*/
			err = errors.New("error dividing numbers")
		}
	}()
	result = divide(x, y)
	return
}

// 3rd party api
func divide(x, y int) int {
	//programmatically creating a panic
	if y == 0 {
		panic(DivideByZeroError)
	}
	result := x / y
	fmt.Println("returning the divide result")
	return result
}
