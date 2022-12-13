package main

import "fmt"

type Employee struct {
	Id     int
	Name   string
	Salary float32
}

func main() {
	// var emp Employee
	// fmt.Println(emp)
	// fmt.Printf("%#v\n", emp)
	// fmt.Printf("%+v\n", emp)
	// var emp Employee = Employee{100, "Magesh", 10000}
	// var emp Employee = Employee{Id: 100, Salary: 10000, Name: "Magesh"}
	/*
		var emp Employee = Employee{
			Id:     100,
			Salary: 10000,
			Name:   "Magesh",
		}
	*/
	emp := Employee{
		Id:     100,
		Salary: 10000,
		Name:   "Magesh",
	}
	fmt.Printf("%+v\n", emp)
	PrintEmployee(emp)
	/*
		e2 := emp
		e2.Id = 101
		fmt.Println(emp, e2)
	*/
	e2 := Employee{
		Id:     100,
		Name:   "Magesh",
		Salary: 10000,
	}
	fmt.Println(emp == e2)

	fmt.Println("After awarding 5000 bonus")
	AwardBonus(&emp, 5000)
	PrintEmployee(emp)

	/* Accessing the fields using the pointer */
	empPtr := &emp
	fmt.Println(empPtr.Name, empPtr.Id, empPtr.Salary)

	//Pointer to a struct
	// x := &Employee{}
	x := new(Employee)
	fmt.Println(x)
}

func PrintEmployee(emp Employee) {
	fmt.Printf("Id = %d, Name = %q, Salary = %.2f\n", emp.Id, emp.Name, emp.Salary)
}

func AwardBonus(emp *Employee, bonus float32) {
	// (*emp).Salary += bonus
	emp.Salary += bonus
}
