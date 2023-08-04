package main

import f"fmt"
import s"strings"
import h"go003Basics/helper"

func main() {
	f.Println(s.ToUpper("Structs:"))
	h.Sep()
	
	var person = struct {
		Id int
		Name string
	}{
		Id: 1,
		Name: "Iaroslav Os",
	}

	f.Printf("Person's ID: %d, NAME: %s\n", person.Id, person.Name)
	h.Sep()
	
	var person2 TPerson;
	person2.Id = 1
	person2.Name = "Iaroslav Os"
	person2.Age = 33
	
	f.Printf("Person's ID: %d, NAME: %s, AGE: %d\n", person2.Id, person2.Name, person2.Age)
	h.Sep()
	
	var person3 = TPerson {
		Id: 2,
		Name: "Comaradeos",
		Age: 24,
	}
	f.Printf("Person's ID: %d, NAME: %s, AGE: %d\n", person3.Id, person3.Name, person3.Age)
	h.Sep()
	
	var person4 = newPerson(1, "Inoqos", 29)
	f.Printf("Person's ID: %d, NAME: %s, AGE: %d\n", person4.Id, person4.Name, person4.Age)

	h.Sep()
	f.Println(person3.getInfo()) // method
	
	h.Sep()
	var pPerson3 *TPerson = &person3 // pointer
	setName(pPerson3, "NewName")
	f.Println(person3.getInfo())
}

type TPerson struct {
	Id int
	Name string
	Age int
}

func newPerson(id int, name string, age int) TPerson {
	return TPerson {
		Id: id,
		Name: name,
		Age: age,
	}
}

// declare method
func (p TPerson) getInfo() string {
	return f.Sprintf("Person's ID: %d, NAME: %s, AGE: %d\n", p.Id, p.Name, p.Age)
}

func setName(p *TPerson, name string) {
	p.Name = name
}