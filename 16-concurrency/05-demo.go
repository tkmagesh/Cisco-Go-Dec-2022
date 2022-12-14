package main

import (
	"fmt"
	"sync"
)

var counter int

func main() {
	wg := &sync.WaitGroup{}
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go fn(wg)
	}
	wg.Wait()
	fmt.Println("Counter :", counter)
}

func fn(wg *sync.WaitGroup) {
	defer wg.Done()
	counter++
}
