package main

import (
	"fmt"
	// "math"
)

const pi = 3.1415

func calculateCircleArea(radius int) float32 {
	return float32(radius) * float32(radius) * pi
}

func printCircleArea(radius int) {
	var area float32 = calculateCircleArea(radius)
	fmt.Printf("Circle area: %.3f\n", area)
}

func main() {
	fmt.Println("funcs")
	printCircleArea(2)
	printCircleArea(3)
	printCircleArea(4)
}