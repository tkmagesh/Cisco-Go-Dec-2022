package calculator

import "fmt"

var op_count int

var ProductRanks map[string]int

func init() {
	ProductRanks = make(map[string]int)
	fmt.Println("calculator[calc.go] initialized")
}

func OpCount() int {
	return op_count
}
