package main

import (
	"fmt"
)

func main() {
	defer catch("попытка 1")
	result := div(2, 0)

	defer catch("попытка 2")
	result = div(2, 0)
	
	fmt.Println(result)
}

func div(a, b int) int {
	return a / b
}

func catch(msg string) {
	if r := recover(); r != nil {
		fmt.Printf("--->>> ПАНИКА: %s (%s)\n", r, msg)
	}
}