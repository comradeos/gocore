package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age int
}

func GetStruct1(p *Person) {
	p.Age = 123
	p.Name = "aaaa"
}

func GetStruct2() Person {
	return Person{
		Name: "Bbbb",
		Age: 123,
	}
}

func Add5Years(p *Person){
	p.Age += 5
}

func main() {
	// p1 := &Person{}
	// GetStruct1(p1)
	// fmt.Println(p1.Name)
	// var p2 = GetStruct2()
	// fmt.Println(p2.Name)

	// var num int = 10
	// var pNum = &num

	// *pNum = 11
	// fmt.Println(num) // value

	// var p = Person{
	// 	Name: "Max",
	// 	Age: 100,
	// }

	// var pP = &p

	// fmt.Println(p.Age)

	// (*pP).Age = 2222

	// fmt.Println(p.Age)

	// Add5Years(pP)

	// fmt.Println(p.Age)

	var num int = 1
	var pNum = &num
	var ppNum = &pNum
	var pppNum = &ppNum

	fmt.Println(num) // 1

	*pNum += 1

	fmt.Println(num) // 2

	**ppNum += 1

	fmt.Println(num) // 3

	***pppNum += 1

	fmt.Println(num) // 4

	



}