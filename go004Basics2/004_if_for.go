package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello");
	const a int = 10
	fmt.Println(a)

	const b int = 10
	b = 25
	fmt.Println(b)

}