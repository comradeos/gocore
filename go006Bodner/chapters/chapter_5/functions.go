package chapter_5

import (
	"fmt"
	"go006Bodner/chapters/helper"
)

func Functions() {
	helper.Line()
	fmt.Println("Functions")

	res := div(5, 2)
	fmt.Println(res)
}

func div(nominator int, denominator int) int {
	if denominator == 0 {
		return 0
	}

	return nominator / denominator
}
