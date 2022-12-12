package main

import "fmt"

func main() {
	var i int32 = 100
	var f float64
	f = float64(i) // use the typename as a function for type conversion
	fmt.Println(f)
}
