package chapter_4

import (
	"fmt"
	"go006Bodner/chapters/helper"
)

func Switches() {
	bEx1()
	helper.Line()
}

func bEx1() {
	words := []string{
		"One",
		"Second",
		"GolangSeemsCool",
		"Yo",
		"A",
	}

	for _, word := range words {
		switch size := len(word); size {
		case 2, 3:
			fmt.Println(word, "it's a short word")
		case 4, 5, 6:
			fmt.Println(word, "it's mid size word")
		case 12, 13, 14, 15, 16:
			fmt.Println(word, "it's COOL long word")
		default:
			fmt.Println("default here")
		}
	}
}
