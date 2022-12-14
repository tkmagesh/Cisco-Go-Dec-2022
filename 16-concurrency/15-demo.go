package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(8)
	ch := genNumbers()
	for data := range ch {
		fmt.Println(data)
	}
}

//producer
func genNumbers() <-chan int {
	count := rand.Intn(20)
	fmt.Println("count :", count)
	ch := make(chan int)
	go func() {
		for i := 1; i <= count; i++ {
			time.Sleep(500 * time.Millisecond)
			ch <- i * 10
		}
		close(ch)
	}()
	return ch
}
