package chapter_4

import (
	"fmt"
	"go006Bodner/chapters/helper"
)

func Shadowing() {
	fmt.Println("Shadowing")

	example1()
	helper.Line()

	example2()
	helper.Line()

	example3()
	helper.Line()

	example4()
	helper.Line()

}

func example4() {
	fmt.Println(true)
	true := 10
	fmt.Println(true)
}

func example3() {
	x := 12
	fmt.Println(x)
	//fmt := "oops"
	//fmt.Println(fmt)
	// fmt.Println undefined (type string has no field or method Println)
}

func example2() {
	x := 10
	if x > 5 {
		fmt.Println(x)
		x, y := 5, 20
		fmt.Println(x, y)
	}
	fmt.Println(x)
}

func example1() {
	x := 10

	if x > 5 {
		fmt.Println(x)
		x := 5
		fmt.Println(x)
	}

	fmt.Println(x)
}
