package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start")

	res1 := make(chan int)
	res2 := make(chan int)
	res3 := make(chan int)

	go sum(2, 2, res1)
	time.Sleep(1 * time.Second)

	go sum(3, 6, res1)
	time.Sleep(1 * time.Second)

	go multiply(7, 7, res2)
	time.Sleep(1 * time.Second)

	go divide(9, 3, res3)
	time.Sleep(5 * time.Second)
}

func sum(a, b int, result chan int) {
	fmt.Printf("a + b = %d, ", (a + b))
}

func multiply(a, b int, result chan int) {
	fmt.Printf("a * b = %d, ", (a * b))
}

func divide(a, b int, result chan int) {
	fmt.Printf("a / b = %d", (a / b))
}
