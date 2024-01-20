package main

import (
	"fmt"
)

func main() {
	deferTest()
}

func deferTest() {
	var a int = 0

	// stack LIFO (Last In First Out)
	defer fmt.Println("defer -> ", a)
	defer fmt.Println("a")
	defer fmt.Println("b")
	defer fmt.Println("c")

	a++
	a++
	a++

	fmt.Println("a =", a)
}