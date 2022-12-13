package main

import (
	"fmt"

	"github.com/fatih/color"

	"github.com/tkmagesh/cisco-go-dec-2022/11-modularity-app/calculator"
	myutils "github.com/tkmagesh/cisco-go-dec-2022/11-modularity-app/calculator/utils"
	_ "github.com/tkmagesh/cisco-go-dec-2022/11-modularity-app/core"
)

func init() {
	fmt.Println("main initialized")
}

func main() {
	fmt.Println("modularity demo app invoked")
	fmt.Println(calculator.Add(100, 200))
	fmt.Println(calculator.Subtract(100, 200))
	fmt.Println("Op Count :", calculator.OpCount())
	color.Red(fmt.Sprintf("is 21 an even number ? %v\n", myutils.IsEven(21)))
	fmt.Printf("is 21 an even number ? %v\n", myutils.IsEven(21))
}
