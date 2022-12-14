package main

import (
	"fmt"
	"time"
)

/*
func main() {
	ch := make(chan int)
	count := 20
	go genNumbers(count, ch)
	for i := 1; i <= count; i++ {
		fmt.Println(<-ch)
	}
}

func genNumbers(count int, ch chan int) {
	for i := 1; i <= count; i++ {
		time.Sleep(500 * time.Millisecond)
		ch <- i * 10
	}
}
*/

func main() {
	count := 20
	ch := genNumbers(count)
	for i := 1; i <= count; i++ {
		fmt.Println(<-ch)
	}
}

func genNumbers(count int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 1; i <= count; i++ {
			time.Sleep(500 * time.Millisecond)
			ch <- i * 10
		}
	}()
	return ch
}
