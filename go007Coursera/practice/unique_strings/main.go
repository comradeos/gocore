package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	fmt.Println("--- Unique Strings App ---")

	var prev string

	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		line := input.Text()

		if line < prev {
			panic("not sorted")
		}

		if line == prev {
			continue
		}

		prev = line

		fmt.Println(line)
	}
}