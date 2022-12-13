package calculator

import "fmt"

func init() {
	fmt.Println("calculator[add.go] initialized")
}

func Add(x, y int) int {
	op_count++
	return x + y
}
