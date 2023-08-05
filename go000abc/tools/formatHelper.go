package tools

import f "fmt"

func Line(size int) {
	if size == 0 {
		size = 50
	}
	for i := 1; i <= size; i++ {
		f.Print("-")
	}
	f.Println()
}
