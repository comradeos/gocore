package main

import "fmt"

func main() {
	var phrase string = "Hello, Go!"
	fmt.Println(phrase)
	
	var i, j, k = 2, 4.4, true
	fmt.Println(i, j, k)

	fmt.Println("----------")
	var myType = MyType { Id:1, Name:"aa" }
	fmt.Printf("myType.Name = %s\n", myType.Name)

} 

type MyType struct {
	Id int
	Name string
}