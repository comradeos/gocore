package chapter2

import (
	"fmt"
)

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

	var r rune = 'æ'
	fmt.Println("r:", r)
	fmt.Printf("r = %c\n", r)

	fmt.Println("Complex")

	var c1 complex64 = complex(2.5, 2.1)
	var c2 complex64 = complex(1.1, 1.2)
	fmt.Println(c1 - c2)

	const myConst1 int = 123
	//myConst1 = 222 //err
	fmt.Println(myConst1)

	var myIntArr1 [3]int
	myIntArr1[0] = 1
	fmt.Println(myIntArr1[0])
	fmt.Println(myIntArr1[1])
	fmt.Println(myIntArr1[2])

	var myIntArr2 = [3]int{1, 2, 3}
	fmt.Println(myIntArr2)

	var myIntArr3 = [10]int{1, 1: 7, 9: 10}
	fmt.Println(myIntArr3)

	var myIntArr4 = [...]int{1, 2, 3}
	fmt.Println(myIntArr4)

	//Slices
	var mySlice1 = []int{1, 2, 3}
	fmt.Println(mySlice1)

	var mySlice2 = make([]int, 2, 4)
	fmt.Println(mySlice2)

	mySlice2[0] = 25
	mySlice2[1] = 50
	mySlice2 = append(mySlice2, 100)

	fmt.Println(mySlice2)

	var mySlice3 = []int{1, 2, 3, 4, 5}
	r1 := mySlice3[1:3]
	fmt.Println(r1)
	r1 = mySlice3[:4]
	fmt.Println(r1)
	r1 = mySlice3[2:]
	fmt.Println(r1)

	var mySlice4 = []int{1, 2, 3, 4, 5}
	r2 := mySlice4[0:3]   // [1 2 3]
	r2[0] = 100           // изменение влияет на исходный срез
	fmt.Println(mySlice4) // [100 2 3 4 5]

	// строки
	var myString1 = "Hello"
	fmt.Println(myString1)

	var sub1 = myString1[0:2]
	sub1 = "a" + sub1[1:]

	var myString2 string = "Hello ☀️"
	fmt.Println(myString2)
	var myString2Bytes = []byte(myString2)
	fmt.Println(myString2Bytes)
	var myString2Runes = []rune(myString2)
	fmt.Println(myString2Runes)

	// maps
	var myMap1 = make(map[string]int)
	myMap1["a"] = 123
	myMap1["b"] = 526
	fmt.Println(myMap1)
	fmt.Println(myMap1["a"])

}
