package calculator

import "fmt"

func init() {
	fmt.Println("calculator[subtract.go] initialized")
}

func Subtract(x, y int) int {
	op_count++
	return x - y
}
