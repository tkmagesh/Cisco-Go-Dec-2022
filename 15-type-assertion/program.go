package main

import "fmt"

type Employee struct {
	Name string
}

func main() {
	// var x interface{}
	var x any
	x = 100
	x = "Tempor ex dolor Lorem minim sunt pariatur tempor incididunt ullamco cupidatat adipisicing. Mollit exercitation ullamco enim irure anim magna reprehenderit culpa. Aute ullamco mollit amet labore incididunt eu amet anim aliquip consequat amet id anim fugiat."
	x = true
	x = 10.00
	x = struct{}{}
	x = Employee{}
	fmt.Println(x)

	x = "Nostrud eiusmod irure dolor proident sunt magna irure exercitation sint dolore Lorem non. Ipsum labore cupidatat voluptate nisi ex ipsum. Incididunt aliquip et ea ullamco. Quis amet occaecat consectetur Lorem sunt non ullamco ad non id. Veniam Lorem esse consequat cillum magna. Consectetur Lorem aliqua proident Lorem labore culpa consequat commodo. Consectetur minim est ipsum commodo cillum sunt cillum qui officia consectetur consectetur id ipsum."
	// unsafe
	// y := x.(int) + 300

	// safe
	x = 200
	if val, ok := x.(int); ok {
		y := val + 300
		fmt.Println(y)
	} else {
		fmt.Println("x is not an int")
	}

	// x = 200
	// x = 99.99
	// x = true
	// x = struct{}{}
	// x = Employee{Name: "Magesh"}
	x = struct{ Id int }{Id: 100}
	switch val := x.(type) {
	case int:
		fmt.Println("x is an int, x + 300 =", val+300)
	case float64:
		fmt.Println("x is a float, 90% of x is =", val*0.9)
	case bool:
		fmt.Println("x is a bool, !x =", !val)
	case struct{}:
		fmt.Println("x is a struct{}")
	case Employee:
		fmt.Println("x is an Employee, Name = ", val.Name)
	default:
		fmt.Println("Unknown type")
	}

}
