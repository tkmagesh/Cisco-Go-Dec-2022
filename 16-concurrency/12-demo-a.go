package main

import (
	"fmt"
)

//Share memory by communicating

//consumer
func main() {
	/*
		var ch chan int
		ch = make(chan int)
	*/
	ch := make(chan int)
	go add(100, 200, ch)
	result := <-ch
	fmt.Println(result)
}

// producer
func add(x, y int, ch chan int) {
	result := x + y
	ch <- result
	fmt.Println("add completed")
}
