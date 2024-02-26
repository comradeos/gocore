package main

import (
	"fmt"
	"ApplicationABC/packages/pNumOne"
	"ApplicationABC/packages/pNumTwo"
	"ApplicationABC/controllers"
)

func main() {
	fmt.Println("ApplicationABC")
	pNumOne.FunctionOne()
	pNumTwo.FunctionTwo()

	a := controllers.GeneralController()
	fmt.Println(a)
}
