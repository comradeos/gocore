package main

import f"fmt"
import h"go003Basics/helper"
import "reflect"

type Employee struct {
	Id int
	Name string
	Age int
	Sex string
	Salary float64
}

func newEmployee(id int, name string, age int, sex string, salary float64) Employee {
	return Employee {
		Id: id,
		Name: name,
		Age: age,
		Sex: sex,
		Salary: salary,
	}
}

func (e Employee) getInfo() {
	f.Printf("Employee info:\n\tId: %d\n\tName: %s\n\tAge: %d\n\tSex: %s\n\tSalary: %.2f\n", e.Id, e.Name, e.Age, e.Sex, e.Salary)
}

func (e * Employee) setSalary(salary float64) {
	e.Salary = salary
}

func main() {
	f.Println("Methods:")
	h.Sep()
	var e1 = newEmployee(1, "Iaroslav Os", 33, "male", 1500)
	e1.getInfo()
	
	h.Sep()
	
	e1.setSalary(1700.00)
	e1.getInfo()

	var a int = 15
	f.Printf("Type: %T\n", a)	
	f.Printf("Type: %v\n", reflect.TypeOf(a))
	f.Printf("Type: %T\n", e1)
	f.Printf("Type: %v\n", reflect.TypeOf(e1))
}