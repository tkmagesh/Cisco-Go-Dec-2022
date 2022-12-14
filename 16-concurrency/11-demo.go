package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		fmt.Println("attempting to send data")
		ch <- 100
		fmt.Println("succeeded sending data")
	}()
	time.Sleep(2 * time.Second)
	result := <-ch
	fmt.Println(result)
}
