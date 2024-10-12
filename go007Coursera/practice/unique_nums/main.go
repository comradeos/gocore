package main

import (
	"fmt"
	"unique_nums/packages"
)

func main() {
	myList := []int { 1, 1, 2, 2, 3, 3 }

	res1 := utils.ContainsInt(myList, 133)
	fmt.Println(res1)

	res2 := utils.GetUniqueNumbers(myList)
	fmt.Println(res2)

	res3 := ReturnTheSame(25)
	fmt.Println(res3)
}

func ReturnTheSame(value int) int {
	return value
}