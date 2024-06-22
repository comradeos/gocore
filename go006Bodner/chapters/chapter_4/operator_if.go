package chapter_4

import (
	"fmt"
	"go006Bodner/chapters/helper"
	"math/rand/v2"
)

func OperatorIf() {
	ex1()
	helper.Line()

}

func ex1() {
	num := rand.IntN(10)

	if num == 0 {
		fmt.Println("That's too low")
	} else if num > 5 {
		fmt.Println("That's too big")
	} else {
		fmt.Println("That's a good number", num)
	}
}
