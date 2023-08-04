package main

import f"fmt"
import h"go003Basics/helper"
import "reflect"

type age int

func (a age) isAdult() bool {
	return a >= 18
}

func main() {
	f.Println("Custom types:")
	h.Sep()

	var a age = age(25)
	
	f.Println(a)
	f.Println(reflect.TypeOf(a))
	f.Println(a.isAdult())
	
	h.Sep()

	var b int = 27
	f.Println(b)
	f.Println(reflect.TypeOf(b))
	// f.Println(b.isAdult())
}