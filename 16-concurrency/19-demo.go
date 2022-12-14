package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		time.Sleep(3 * time.Second)
		ch1 <- 100
	}()

	go func() {
		time.Sleep(4 * time.Second)
		data3 := <-ch3
		fmt.Println("data3 :", data3)
	}()

	go func() {
		time.Sleep(5 * time.Second)
		ch2 <- 200
	}()

	for i := 0; i < 3; i++ {
		select {
		case data1 := <-ch1:
			fmt.Println(data1)
		case data2 := <-ch2:
			fmt.Println(data2)
		case ch3 <- 300:
			fmt.Println("Data sent to channel 3")
			/* default:
			fmt.Println("No channer operations succeeded") */
		}
	}
}
