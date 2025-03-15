package main

import "fmt"

func variables() {
	sep()
	red("01 variables")
	sep()

	var v1 string = "Hello"
	v2 := "World"
	var v3 = "!"
	fmt.Println(v1, v2, v3)
}
