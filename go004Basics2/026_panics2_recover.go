package main

import (
	"fmt"
)

func main() {
	fmt.Println("Start")

	res := div(10, 0)
	_ = res	
	
	fmt.Println("End")
}

func div(a, b int) int {
	defer func(){
		if err := recover(); err != nil {
			fmt.Printf("Panic happened: %s\n", err)
		}
	}()
	return a / b
}
