package modules

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	Salary int
}

func GoInheritance() {
	employee := Employee{
		Person: Person{
			Name: "Name1",
			Age:  24,
		},
		Salary: 2500,
	}

	fmt.Println(employee.Name)
	fmt.Println(employee.Age)
	fmt.Println(employee.Salary)
}
