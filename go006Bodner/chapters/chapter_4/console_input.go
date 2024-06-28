package chapter_4

import "fmt"

func ConsoleInput() {
	var num int

	fmt.Print("Enter a number: ")
	_, err := fmt.Scan(&num)

	if err != nil {
		fmt.Println("Error...")
	}

	fmt.Printf("num = %d\n", num)
}
