package main

import (
	"errors"
	"fmt"
	mp "go000abc/myPackage"
	"math"
)

func main() {
	var powResult float64 = math.Pow(3, 5)
	fmt.Printf("powres: %.3f\n", powResult)

	var err error = errors.New("ERRR")

	if err != nil {
		fmt.Println(err.Error())
	}

	var num int = 2314
	pNum := &num
	fmt.Println(pNum)
	fmt.Println(mp.MyNumber)

	mp.Sep()

	fmt.Println(divZero(12, 2))
	fmt.Println(divZero(12, 0))

	//fmt.Println(5.0 / 2.0)

}

// try catch example
func divZero(a int, b int) int {
	defer func() {
		if r := recover(); r != nil { // catch
			fmt.Println("шось не так - шось не ділеться")
		}
	}()
	return a / b // try
}

/*

go build main.go && ./main

*/
