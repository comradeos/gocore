package main

import (
	"fmt"
)

type Person struct { // структура (типа класса)
	name string
	age int
	X, Y int
}

func (p *Person) info() { // метод
	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
	fmt.Println("X:", p.X)
	fmt.Println("Y:", p.Y)
}

func main() {
	var p1 Person
	
	p1 = Person {
		name: "Ivan",
		age: 30,
		X: 10,
		Y: 20,
	}


	fmt.Println(p1)

	p2 := Person {
		age: 3131,
	}

	fmt.Println(p2)

	p1.info()

}