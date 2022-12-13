/* deferred functions */
package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("	deferred - main")
	}()
	fmt.Println("main started")
	f1Result := f1(2)
	fmt.Println("f1Result :", f1Result)
	fmt.Println("main completed")
}

func f1(x int) (result int) {
	/*
		defer func() {
			fmt.Println("	deferred - f1[1]")
		}()
		defer func() {
			fmt.Println("	deferred - f1[2]")
		}()
		defer func() {
			fmt.Println("	deferred - f1[3]")
		}()
	*/

	defer fmt.Println("	deferred - f1[1]")
	defer fmt.Println("	deferred - f1[2]")
	defer fmt.Println("	deferred - f1[3]")
	defer func() {
		defer fmt.Println("	deferred - f1[4]")
		result = x * 1000
	}()
	fmt.Println("f1 started")
	f2()
	fmt.Println("f1 completed")
	result = 100
	return
}

func f2() {
	defer func() {
		fmt.Println("	deferred - f2")
	}()
	fmt.Println("f2 started")
	fmt.Println("f2 completed")
}
