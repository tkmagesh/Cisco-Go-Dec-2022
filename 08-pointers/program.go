/* pointers */
package main

import "fmt"

func main() {
	/*
		var no int = 100
		var noPtr *int = &no
	*/
	no := 100
	noPtr := &no
	fmt.Println(no, noPtr)

	//dereferencing - pointer -> value
	var val int = *noPtr
	fmt.Println(val)

	fmt.Println("Before incrementing, no :", no)
	increment(&no)
	fmt.Println("After incrementing, no :", no)

	n1, n2 := 10, 20
	fmt.Println("before swapping, n1, n2 : ", n1, n2)
	swap(&n1, &n2)
	fmt.Println("after swapping, n1, n2 : ", n1, n2)
}

func increment(x *int) {
	*x++
}

func swap(x, y *int) {
	*x, *y = *y, *x
}
