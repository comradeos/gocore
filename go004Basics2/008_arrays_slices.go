package main

import (
	"fmt"
)

func main() {
	var array [5] int // массив из 5 элементов типа int
	
	array[0] = 10
	array[1] = 20
	array[4] = 30	

	fmt.Println(array)

	var array2 [5] int = [...] /* посчитать */ int {1, 2, 3, 4, 5}
	fmt.Println(array2)


	var slice1 [] int
	slice1 = append(slice1, 10)
	slice1 = append(slice1, 20)
	slice1 = append(slice1, 30)
	slice1 = append(slice1, 40)
	fmt.Println(slice1)

	slice1 = append(slice1[:0], slice1[1:]...) // удаление элемента
	fmt.Println(slice1)

	fmt.Println(len(slice1)) // длина
	fmt.Println(cap(slice1)) // вместимость


	fmt.Println()
	fmt.Println()
	fmt.Println()

	var mySlice [] int
	
	mySlice = append(mySlice, 10) ; mySliceLen := len(mySlice) ; mySliceCap := cap(mySlice)
	mySlice = append(mySlice, 10, 10, 10, 10, 10, 10, 10, 10, 10) ; mySliceLen = len(mySlice) ; mySliceCap = cap(mySlice)
	mySlice = append(mySlice, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10) ; mySliceLen = len(mySlice) ; mySliceCap = cap(mySlice)
	mySlice = append(mySlice, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10) ; mySliceLen = len(mySlice) ; mySliceCap = cap(mySlice)
	
	fmt.Println(mySlice, "Len =",  mySliceLen,  "Cap =", mySliceCap)
	
}