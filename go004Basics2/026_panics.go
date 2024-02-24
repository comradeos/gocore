package main

import (
	"fmt"
)

func main() {
	var arr1 []int
	arr1 = []int {
		1, 
		2, 
		3,
	}

	fmt.Println(arr1[0])
	// fmt.Println(arr1[1])
	// fmt.Println(arr1[2])

	p1 := &Person{"John", 30}
	p1 = nil
	// p1.Info()
	p1.Hello()

	panic("This is a panic")
}

type Person struct {
	Name string
	Age int
}

func (p *Person) Info() {
	fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)
}

func (p *Person) Hello() {
	fmt.Println("ok")
}
