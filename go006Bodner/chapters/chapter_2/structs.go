package chapter_2

import (
	"fmt"
	"go006Bodner/chapters/helper"
)

func Structs() {
	helper.Line()
	fmt.Println("Structs:")

	type firstPerson struct {
		age  int
		name string
	}

	type secondPerson struct {
		age  int
		name string
	}

	var v1p1 firstPerson
	v1p1.name = "A"
	v1p1.age = 1

	var v2p1 secondPerson
	v2p1.name = "a"

}
