package main

import "fmt"

type Employee struct {
	Id     int
	Name   string
	Salary float32
	Org    *Organization
}

type Organization struct {
	Name string
	City string
}

func main() {
	cisco := &Organization{
		Name: "Cisco",
		City: "Bengaluru",
	}
	e1 := Employee{
		Id:     100,
		Name:   "Magesh",
		Salary: 10000,
		Org:    cisco,
	}

	e2 := Employee{
		Id:     101,
		Name:   "Suresh",
		Salary: 20000,
		Org:    cisco,
	}

	cisco.City = "Pune"
	fmt.Println(e1.Org)
	fmt.Println(e2.Org)

}
