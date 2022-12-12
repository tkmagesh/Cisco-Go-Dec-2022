package main

import "fmt"

func main() {
	var c1 complex64
	c1 = 4 + 5i
	fmt.Println(c1)
	fmt.Printf("real : %v\n", real(c1))
	fmt.Printf("imaginary : %v\n", imag(c1))
	var c2 complex64 = 9 + 7i
	result := c1 + c2
	fmt.Println(result)
}
