package helper

import "fmt"

func Line() {
	for i := 0; i < 30; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}
