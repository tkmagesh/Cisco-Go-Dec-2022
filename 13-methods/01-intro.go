package main

import "fmt"

type Employee struct {
	Id     int
	Name   string
	Salary float32
}

func (emp Employee) WhoAmI() {
	fmt.Println("I am ", emp.Name)
}

func (emp Employee) PrintEmployee() {
	fmt.Printf("Id = %d, Name = %q, Salary = %.2f\n", emp.Id, emp.Name, emp.Salary)
}

func (emp *Employee) AwardBonus(bonus float32) {
	// (*emp).Salary += bonus
	emp.Salary += bonus
}

func main() {
	/*
		emp := Employee{
			Id:     100,
			Name:   "Magesh",
			Salary: 10000,
		}
		emp.WhoAmI()
		emp.PrintEmployee()
		emp.AwardBonus(10000)
		emp.PrintEmployee()
	*/

	empPtr := &Employee{
		Id:     100,
		Name:   "Magesh",
		Salary: 10000,
	}
	empPtr.WhoAmI()
	empPtr.PrintEmployee()
	empPtr.AwardBonus(10000)
	empPtr.PrintEmployee()
}
