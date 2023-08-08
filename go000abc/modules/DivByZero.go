package modules

import "fmt"

// try catch example

func divZero(a int, b int) int {
	defer func() {
		if r := recover(); r != nil { // catch
			fmt.Println(r)
			fmt.Println("шось не так - шось не ділеться")
		}
	}()
	return a / b // try
}

func TryExceptPanicRecover() {
	fmt.Println(divZero(4, 2))
	fmt.Println(divZero(4, 0))
}
