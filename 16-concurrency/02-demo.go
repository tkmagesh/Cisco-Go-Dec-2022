package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("main started")
	wg.Add(1) // increment the counter by 1
	go f1()   //scheduling f1() for execution
	f2()
	wg.Wait() // block until the counter becomes 0
	fmt.Println("main completed")
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(5 * time.Second)
	fmt.Println("f1 completed")
	wg.Done()
}

func f2() {
	fmt.Println("f2 invoked")
}
