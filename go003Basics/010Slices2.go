package main

import (
	"fmt"
	h "go003Basics/helper"
)

func main() {
	fmt.Println("Slices2:")
	h.Sep()

	var slices = []string {
		"a",
		"b",
		"c",
		"d",
	}

	var res []string;
	res = slices[1:3] // включ:невключ

	fmt.Println(res)
	res[0] = "-1-"
	res[1] = "-2-"

	h.Sep()

	fmt.Println(slices)

}