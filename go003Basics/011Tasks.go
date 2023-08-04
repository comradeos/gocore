package main

import "fmt"
import "math"
import h "go003Basics/helper"

func myPow(a int, b int) int {
	var result int = 1;
	if b <= 0 { return result }
	for i:=1; i<=b; i++ {
		result *= a
	}
	return result
}

func pow2To(n int) {
	if n <= 0 { n = 1 }
	for i:=1; i<=n; i++ {
		powResult := math.Pow(2, float64(i))
		fmt.Println(int(powResult))
	}
}

func main() {
	fmt.Println("Tasks")
	h.Sep()
	pow2To(4)
	h.Sep()
	fmt.Println(myPow(2,4))
}