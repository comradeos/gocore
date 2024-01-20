package main

import (
	"fmt"
	// "time"
)

func main() {
	for i := 0; i < 10; i += 2 { // 0 2 4 6 8 
		fmt.Printf("%d ", i)
	}

	fmt.Println()

	var sum int = 0

	for sum < 10 { // 5 10
		sum += 5
		fmt.Print(sum, " ")
	}

	fmt.Println()

	var number int = 0

	// for { бесконечный цикл
	// 	number += 500
	// 	fmt.Print(number, " ")
	// 	time.Sleep(500 * time.Millisecond)
	// }
	
	if number > 0 {
		fmt.Println("number > 0")
	} else {
		fmt.Println("number <= 0")
	}

	if a := test(number); a == -1 {
		fmt.Println("a == -1")
	}

	switch number {
		case 0:
			fmt.Println("number == 0")
		case 1:
			fmt.Println("number == 1")
		default:
			fmt.Println("number != 0 && number != 1")
	}
}

func test(number int) int {
	if number > 0 {
		return 1
	}
	return -1
}