package tests

import "fmt"

type MyStruct1 struct {
	Name string
}

type MyStruct2 struct {
	Name string
}

func ST002NilsEquation() {
	var struct1 MyStruct1
	var struct2 MyStruct2

	fmt.Println(struct1)
	fmt.Println(struct2)

	//fmt.Println(struct1 == struct2) // cannot compare struct1 == struct2 (mismatched types MyStruct1 and MyStruct2)
	fmt.Println(struct1 == (MyStruct1)(struct2)) // true

	//struct1 = (MyStruct1)(struct2)

	var pStruct1 *MyStruct1
	var pStruct2 *MyStruct2
	fmt.Println(pStruct1)
	fmt.Println(pStruct2)
	//fmt.Println(pStruct1 == pStruct2)             // error
	fmt.Println(pStruct1 == (*MyStruct1)(pStruct2)) // true
}
