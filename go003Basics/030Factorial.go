package main

import (
	"fmt"
	"time"
	// "math/big"
)

func factorial(number int) int {
	if number == 1 {
		return number
	} 
	return number * factorial(number - 1)
}

func main() {
	result := factorial(7)
	fmt.Println(result)
	time.Sleep(time.Millisecond)
}