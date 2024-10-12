package main

import (
	"fmt"
	"bufio"
	"os"
	"io"
)

func main() {
	fmt.Println("--- Unique Strings App 2 ---")

	err := unqi(os.Stdin, os.Stdout)

	if err != nil {
		panic(err)
	}
}

func unqi(input io.Reader, output io.Writer) error {
	var prev string

	in := bufio.NewScanner(input)

	for in.Scan() {
		line := in.Text()

		if line < prev {
			return fmt.Errorf("not sorted")
		}

		if line == prev {
			continue
		}

		prev = line

		fmt.Println(line)
	}

	return nil
}