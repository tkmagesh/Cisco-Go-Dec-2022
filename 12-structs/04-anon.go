package main

import "fmt"

func main() {
	product := struct {
		Id   int
		Name string
		Cost float32
	}{
		Id:   100,
		Name: "Pen",
		Cost: 10,
	}
	fmt.Println(product)
	PrintProduct(product)

	empty := struct{}{}
	fmt.Println(empty)
}

func PrintProduct(p struct {
	Id   int
	Name string
	Cost float32
}) {
	fmt.Println(p.Id, p.Name, p.Cost)
}
