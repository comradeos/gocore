package main 

import (
	"fmt"
	"errors"
	"math"
)

const pi = 3.1415

func main() {
	printCircleArea(2)
	printCircleArea(4)
	printCircleArea(-4)
}

func calculateCircleArea(radius int) (float64, error) {
	if radius < 0 {
		return float64(0), errors.New("Radius can't be negatinve")
	}
	
	var fRadius float64 = float64(radius)
	return math.Pow(fRadius, 2) * pi, nil
}

func printCircleArea(radius int) {
	fmt.Println("-----------------------------")
	fmt.Printf("Circle radius: %d\n", radius)
	fmt.Println("Circle area formula: A=pi*R*R")
	
	area, err := calculateCircleArea(radius)
	
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	
	fmt.Printf("Circle area: %.3f\n", area)
}