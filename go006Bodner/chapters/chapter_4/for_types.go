package chapter_4

import (
	"fmt"
	"go006Bodner/chapters/helper"
)

func ForTypes() {
	ftEx1()
	helper.Line()

	ftEx2()
	helper.Line()

	//ftEx3()
	helper.Line()

	ftEx4()
	helper.Line()

	ftEx5()
	helper.Line()
}

func ftEx5() {
	values := []int{2, 3, 4, 5, 3, 5, 2, 3}

	for i, v := range values {
		fmt.Println(i, v)
	}
}

func ftEx4() {
	// do while
	x := 0
	for {
		fmt.Println(x)
		if x >= 5 {
			break
		}
		x = x + 2
	}
}
func ftEx3() {
	for {
		fmt.Println("hello")
	}
}

func ftEx2() {
	// типа while
	i := 1
	for i < 10 {
		fmt.Print(i, ", ")
		i *= 2
	}
	fmt.Println()
}

func ftEx1() {
	// полный вид
	for i := 1; i < 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()
}
