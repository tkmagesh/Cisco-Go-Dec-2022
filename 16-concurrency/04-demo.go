package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// var wg sync.WaitGroup
	wg := &sync.WaitGroup{}
	noOfGoroutines := flag.Int("count", 0, "Number of goroutines to execute")
	flag.Parse()
	rand.Seed(7)
	fmt.Printf("Starting %d goroutines... Hit ENTER to start\n", *noOfGoroutines)
	fmt.Scanln()
	for i := 1; i <= *noOfGoroutines; i++ {
		wg.Add(1)    // increment the counter by 1
		go fn(i, wg) //scheduling f1() for execution
	}
	wg.Wait() // block until the counter becomes 0
	fmt.Println("Hit ENTER to shutdown!")
	fmt.Scanln()
	fmt.Println("main completed")
}

func fn(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("fn[%d] started\n", id)
	time.Sleep(time.Duration(rand.Intn(20)) * time.Second)
	fmt.Printf("fn[%d] completed\n", id)
}
