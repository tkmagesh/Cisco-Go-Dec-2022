package main

import "fmt"

func main() {
	y := 200
	f1(y)
}

//go:noinline
func f1(no int) {
	x := 100
	fmt.Println(x)
	/* func() {
		x += 100
	}() */
}
