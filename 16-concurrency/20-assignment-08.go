/*
Modify the program such that it generates fibonocci numbers until the user hits ENTER key
*/

package main

import "fmt"

func main() {
	ch := generateFibonocci()
	for no := range ch {
		fmt.Println(no)
	}
}

func generateFibonocci() <-chan int {
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
