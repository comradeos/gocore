package main

import f "fmt"

type Human interface {
	introduce()
}

type Person struct {
	Name string
	Age int
}

type Employee struct {
	Name string
	Salary float32
}

func (p Person) introduce() {
	f.Printf("Hi! I'm a person! My name is %s!\n", p.Name)
}

func (e Employee) introduce() {
	f.Printf("Hi! I'm an employee! My name is %s!\n", e.Name)
}

func greet(h Human) {
	h.introduce()
} 



func main() {
	f.Println("Ploymorph:")

	var obj1 Human
	
	obj1 = Person {
		Name: "Obj1Person",
	}

	greet(obj1)
	// obj1.introduce()

	obj1 = Employee {
		Name: "Obj2Employee",
	}
	
	greet(obj1)
	// obj1.introduce()
	
}