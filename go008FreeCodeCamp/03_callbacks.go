package main

import "fmt"

func callbacks() {
	sep()
	red("03 callbacks")
	sep()

	someFunction(5, 5, funcToCall)
}

func funcToCall() {
	fmt.Println("funcToCall")
}

func someFunction(a int, b int, callback func()) {
	fmt.Println("a:", a, "b:", b)
	callback()
}
