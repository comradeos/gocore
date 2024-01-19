package main

import (
	"fmt"
)

const myConstA = "A"

func test() (int, int) {
	var a int = 1
	var b int = 2
	return a, b
}

type myStruct struct {
	a int
	b int
	c int
}

func test2() myStruct {
	var a myStruct
	a.a = 1
	a.b = 2
	a.c = 3
	return a
}

func test3() (a, b, c int) {
	a = 1
	b = 2
	c = 3
	return
}

func test4() (int, int, int) {
	a := 1222
	b := 2222
	c := 3222
	return a, b, c
}

func main() {
	fmt.Println(test());
	var s = test2();
	fmt.Println(s.a);
	fmt.Println(test3());
	fmt.Println(test4());
}