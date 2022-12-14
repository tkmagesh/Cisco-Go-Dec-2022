/*
Modify the program such that it generates fibonocci numbers until the user hits ENTER key
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	timeoutCh := timeout(5 * time.Second)
	ch := generateFibonocci(timeoutCh)
	for no := range ch {
		fmt.Println(no)
	}
}

func timeout(d time.Duration) <-chan time.Time {
	stopCh := make(chan time.Time)
	go func() {
		time.Sleep(d)
		stopCh <- time.Now()
	}()
	return stopCh
}

func generateFibonocci(stopCh <-chan time.Time) <-chan int {
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
