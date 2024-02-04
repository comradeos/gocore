package main

import (
	"fmt"
	// "errors"
)

func main() {
	defer func(){
		if r := recover(); r != nil {
			fmt.Println("Recovered from", r)
		}
	}()

	result := divide(12, 0)
	fmt.Println("The result is", result)
}

func divide(a, b int) int {
	return a / b
}