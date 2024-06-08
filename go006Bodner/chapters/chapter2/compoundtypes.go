package chapter2

import "fmt"

type MyStruct struct {
	Id   int
	Name string
}

func CompoundTypes() {
	fmt.Println("Compound types:")

	var pMyStruct1 *MyStruct
	fmt.Println(pMyStruct1)

	//pMyStruct1 = &MyStruct{}
	pMyStruct1 = new(MyStruct)
	pMyStruct1.Name = "Hello World"
	fmt.Println(pMyStruct1)

	Line()

	var myMap = make(map[string]int)
	myMap["a"] = 123
	myMap["b"] = 0
	fmt.Println(myMap)

	v, ok := myMap["a"]
	fmt.Println(v, ok)

	v, ok = myMap["GGG"]
	fmt.Println(v, ok)
}
