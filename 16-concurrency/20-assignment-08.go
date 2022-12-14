/*
Modify the program such that it generates fibonocci numbers until the user hits ENTER key
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := generateFibonocci()
	for no := range ch {
		fmt.Println(no)
	}
}

func generateFibonocci() <-chan int {
	ch := make(chan int)
	stopCh := make(chan struct{})
	go func() {
		fmt.Scanln()
		stopCh <- struct{}{}
	}()

	go func() {
		x, y := 0, 1
	LOOP:
		for {
			select {
			case <-stopCh:
				break LOOP
			case ch <- x:
				time.Sleep(500 * time.Millisecond)
				x, y = y, x+y
			}

		}
		close(ch)
	}()
	return ch
}
