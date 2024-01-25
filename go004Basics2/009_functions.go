package main

import (
	"fmt"
)

func main() {
	// функции c колбэками
	res := doSomething(forCallback, "Hello, world!")
	fmt.Println(res)
	res = doSomething(forCallback2, "Good bye, world!")
	fmt.Println(res)

	fmt.Println()

	var res2 func(int) int = totalPrice(100)

	fmt.Println(res2(10))  // 90 (замыкание)
	fmt.Println(res2(10))  // 80 (замыкание)
	fmt.Println(res2(10))  // 70 (замыкание)

}

// функция для колбэка
func forCallback(a int, b int) int {
	return a + b
}

// функция для колбэка 2
func forCallback2(a int, b int) int {
	return (a * b * 2) * 3
}

// функция с колбэком
func doSomething(callback func(int, int) int, str string) int {
	fmt.Println(str)
	return callback(1, 2)
}

// замыкания
func totalPrice(initPrice int) func(int) int {
	var result int = initPrice
	return func(discount int) int {
		 result -= discount
		 return result
	}
}
