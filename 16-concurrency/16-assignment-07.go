/*
Write a goroutine to generate N fibonocci numbers (N is the input)
Print the generated fibonocci numbers in the main function
*/

package main

import "fmt"

func main() {
	ch := generateFibonocci(10)
	for no := range ch {
		fmt.Println(no)
	}
}

func generateFibonocci(count int) <-chan int {
	ch := make(chan int)
	go func() {
		x, y := 0, 1
		for i := 0; i < count; i++ {
			ch <- x
			x, y = y, x+y
		}
		close(ch)
	}()
	return ch
}
