package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// var wg sync.WaitGroup
	wg := &sync.WaitGroup{}
	fmt.Println("main started")
	wg.Add(1) // increment the counter by 1
	go f1(wg) //scheduling f1() for execution
	f2()
	wg.Wait() // block until the counter becomes 0
	fmt.Println("main completed")
}

func f1(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("f1 started")
	time.Sleep(5 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
