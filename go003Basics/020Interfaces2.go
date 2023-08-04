package main

import f"fmt"
import h"go003Basics/helper"

type IPerson interface {
	intro()
}

type Person struct {
	Name string
	Age int
}

func (p Person) intro() {
	f.Printf("Hello, my name is %s, I'm %d years old!", p.Name, p.Age)
}

func main() {
	h.Sep()
	f.Println("Interfaces:")

	var person IPerson

	person = Person {
		Name: "Ivan",
		Age: 25,
	}
	
	person.intro()
	f.Println(person)
}
