package chpts

import (
	"fmt"
)

func Ch001DataTypes2() {
	fmt.Println("Hello from Ch001DataTypes2()")

	// массив

	var arr1 [5]int // [0 0 0 0 0]
	var arr2 [5]int = [5]int{1, 2, 3, 4, 5}

	fmt.Println("arr1", arr1)
	fmt.Println("arr2", arr2)

	const size = 5
	var arr3 = [size]int{21, 22, 23, 24, 2 + 5}
	fmt.Println("arr3", arr3)

	// опредилить размер массива по количеству элементов
	arr4 := [...]int{31, 32, 33, 34, 35, 36, 37, 38, 39, 40}
	fmt.Println("arr4", arr4)

	// срез

	var slice1 []int = []int{1, 2, 3, 4, 5}

	fmt.Println("slice1", slice1)

	// append

	slice1 = append(slice1, 6)
	fmt.Println("slice1", slice1)

	var slice2 = make([]int, 5, 10)
	fmt.Println("slice2", slice2)

	var slice3 = make([]int, 15)
	fmt.Println("slice3", slice3)

	slice2 = append(slice2, slice3...)
	fmt.Println("slice2", slice2)
}
