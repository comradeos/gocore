package main

import ("fmt")

const pi = 3.1415

func main() {
	fmt.Println("Conditions:")
	printCircleArea(2)
	printCircleArea(4)
	printCircleArea(6)
	printCircleArea(-1)
}

func printCircleArea(radius int) {
	if 2 > 0 && radius < 0 {
		fmt.Println("Error: radius can't be negative!")
		return
	}
	
	var fRadius = float32(radius)
	var area float32 = fRadius * fRadius * pi

	fmt.Printf("Circle area: %.3f\n", area)
}