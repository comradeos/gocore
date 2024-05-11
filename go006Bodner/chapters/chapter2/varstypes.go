package chapter2

import "fmt"

func TypesVars() {
	fmt.Println("Variables, Types")

	var num1 int = 10
	fmt.Println("num1:", num1)

	var num2 int
	fmt.Println("num2:", num2)

	num3 := 20
	fmt.Println("num3:", num3)

	var (
		num4 int = 30
		num5 int = 40
		num6 int
	)

	fmt.Println("num4:", num4)
	fmt.Println("num5:", num5)
	fmt.Println("num6:", num6)

	x, y := 123, "hello"
	fmt.Println("x:", x)
	fmt.Println("y:", y)

	x, a := 333, "dsd"
	fmt.Println("x:", x)
	fmt.Println("a:", a)

	var b *int = nil
	fmt.Println("b:", b)

	const (
		Red = iota
		Green
		Blue
	)

	fmt.Println("Red:", Red)
	fmt.Println("Green:", Green)
	fmt.Println("Blue:", Blue)

	var myArray = [5]int{} // массив
	myArray[0] = 1
	myArray[4] = 1
	fmt.Println("myArray:", myArray)

	var mySlice = []int{1, 2, 3} // срез
	fmt.Println("mySlice:", mySlice)
	mySlice = append(mySlice, 4)
	fmt.Println("mySlice:", mySlice)

	var myMap = map[string]int{"one": 1, "two": 2}
	fmt.Println("myMap:", myMap)
	myMap["three"] = 3
	fmt.Println("myMap:", myMap)

}
