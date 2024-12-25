package main

import (
	"fmt"
	"time"
)

func main() {
	var startTime = time.Now().UnixNano()

	var res int
	res = fibonacci(30)
	
	fmt.Println("res: ", res)

	var endTime int64 = time.Now().UnixNano()

	fmt.Println("Time taken: ", (endTime - startTime) / 1000000, "ms")
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}