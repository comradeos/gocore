package chapter_4

import (
	"fmt"
	"go006Bodner/chapters/helper"
)

func Labels() {
	fmt.Println("Labels")
	helper.Line()

	aEx1()

}

func aEx1() {
	words := []string{
		"one",
		"two",
		"three",
		"four",
	}

outer:
	for _, word := range words {
		for _, r := range word {
			if r == 'h' {
				continue outer
			}
			fmt.Println(r, string(r))
		}
	}
}
