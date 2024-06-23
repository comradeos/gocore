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

	ftEx6()
	helper.Line()

	ftEx7()
	helper.Line()
}

func ftEx7() {
	// обход элементов строки
	samples := []string{
		"hello",
		"apple_π!",
	}

	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
		}
		fmt.Println()
	}
}

func ftEx6() {
	// обход карты может меняться
	m := map[int]string{
		1: "a",
		2: "b",
		3: "c",
	}

	for i := 0; i < 10; i++ {
		for k, v := range m {
			fmt.Print(k, v, " ")
		}
		fmt.Println("*")
	}
}

func ftEx5() {
	// обход среза
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
