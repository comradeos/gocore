package main

import (
	"fmt"
	"strconv"
)

func main() {
	var number int = 10 // число 
	var pNumber *int = &number // указатель на число
	var ppNumber **int = &pNumber // указатель на указатель на число
	var pppNumber ***int = &ppNumber // указатель на указатель на указатель на число

	fmt.Println("number =", number)
	*pNumber = 20 // разыменование указателя
	fmt.Println("number =", number)
	**ppNumber = 30 // разыменование указателя на указатель
	fmt.Println("number =", number)
	***pppNumber = 40 // разыменование указателя на указатель на указатель
	fmt.Println("number =", number)


	var input string
	fmt.Println("Enter a number:")
	fmt.Scan(&input) // вводим число

	_, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("input =", input)
}