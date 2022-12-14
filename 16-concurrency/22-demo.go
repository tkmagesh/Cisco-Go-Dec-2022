/*
Modify the program such that it generates fibonocci numbers until the user hits ENTER key
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	stopCh := getStopSignal()
	ch := generateFibonocci(stopCh)
	for no := range ch {
		fmt.Println(no)
	}
}

func getStopSignal() <-chan struct{} {
	stopCh := make(chan struct{})
	fmt.Println("Hit ENTER to stop...")
	go func() {
		fmt.Scanln()
		stopCh <- struct{}{}
	}()
	return stopCh
}

func generateFibonocci(stopCh <-chan struct{}) <-chan int {
	ch := make(chan int)
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
