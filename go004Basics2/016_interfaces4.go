package main

import (
    "fmt"
)

func main() {
	var a interface{}
	fmt.Printf("%v, %T\n", a, a)
	a = "test"
	fmt.Printf("%v, %T\n", a, a)

	a = 10
	fmt.Printf("%v, %T\n", a, a)

	var s1 interface{} = "abc"
	s2, ok := s1.(int)  // при приведении типа, если тип не совпадает, то второй параметр будет false
	fmt.Println(s2, ok)
}