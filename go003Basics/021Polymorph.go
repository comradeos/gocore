package main

import f"fmt"
import h"go003Basics/helper"
// import r"reflect"

type IPerson interface {
	introduce(nameOnly bool)
}

type Person struct {
	Id int
}

func (p Person) introduce(nameOnly bool) {
	f.Println("Hello")
}

func main() {
	f.Println("Polymorph:")
	h.Sep()

	var p1 IPerson
	p1 = Person {
		Id: 1,
	}

	p1.introduce(true)

	var p2 interface {}
	p2 = Person {
		Id: 2,
	}

	if _, ok := p2.(Person); ok {
		f.Println("1111PERSON")
	}



	// typy check 
	var num interface {}
	num = 25
	if _, ok := num.(int); ok {
		f.Println("this is int")
	}

	/*
	if r.TypeOf(num) == "int" {
		f.Println("THIS IS INT")
	}
	*/
}