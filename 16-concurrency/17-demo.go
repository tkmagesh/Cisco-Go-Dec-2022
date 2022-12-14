package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	oddCh := generateOdd()
	evenCh := generateEven()
	resultCh := fanIn(oddCh, evenCh)
	for data := range resultCh {
		fmt.Println(data)
	}
}

func fanIn(chs ...<-chan int) <-chan int {
	resultCh := make(chan int)
	go func() {
		wg := sync.WaitGroup{}
		for _, ch := range chs {
			wg.Add(1)
			go func(ch <-chan int) {
				for data := range ch {
					resultCh <- data
				}
				wg.Done()
			}(ch)
		}
		wg.Wait()
		close(resultCh)
	}()
	return resultCh
}

func generateOdd() <-chan int {
	ch := make(chan int)
	go func() {
		for i := 1; i < 10; i++ {
			time.Sleep(300 * time.Millisecond)
			ch <- (i * 2) - 1
		}
		close(ch)
	}()
	return ch
}

func generateEven() <-chan int {
	ch := make(chan int)
	go func() {
		for i := 1; i < 10; i++ {
			time.Sleep(500 * time.Millisecond)
			ch <- (i * 2)
		}
		close(ch)
	}()
	return ch
}
