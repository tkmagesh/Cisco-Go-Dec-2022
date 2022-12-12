package main

import (
	"errors"
	"fmt"
)

var DivideByZeroError error = errors.New("divisor cannot be 0")

func main() {
	x, y := 100, 0
	result, err := divide(x, y)
	if err == DivideByZeroError {
		fmt.Println("Do not attempt to divide by zero")
		return
	}
	if err != nil {
		fmt.Println("something went wrong!")
		fmt.Println(err)
		return
	}
	fmt.Printf("Dividing %d by %d, result = %d\n", x, y, result)
}

/*
func divide(x, y int) (int, error) {
	if y == 0 {
		divieByZeroError := fmt.Errorf("cannot divide %d by %d", x, y)
		return 0, divieByZeroError
	}
	return x / y, nil
}
*/

/*
func divide(x, y int) (result int, err error) {
	if y == 0 {
		err = fmt.Errorf("cannot divide %d by %d", x, y)
		return
	}
	result = x / y
	return
}
*/

/*
func divide(x, y int) (result int, err error) {
	if y == 0 {
		err = errors.New("divisor cannot be 0")
		return
	}
	result = x / y
	return
}
*/

func divide(x, y int) (result int, err error) {
	if y == 0 {
		err = DivideByZeroError
		return
	}
	result = x / y
	return
}
