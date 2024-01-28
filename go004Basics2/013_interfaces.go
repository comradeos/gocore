package main

import (
	"fmt"
)

func main() {
	fmt.Println("Interfaces")

	var myInterface MyInterface

	myInterface = &MyStruct1 {1, 2}
	sum := myInterface.Sum()
	fmt.Println(sum)
}

type MyInterface interface {
	Sum() int
	MyMethod1()
}	

type MyStruct1 struct {
	A int
	B int
}

func (s *MyStruct1) Sum() int {
	return s.A + s.B
}

func (s *MyStruct1) MyMethod1() {
    fmt.Println("MyMethod1 called")
}