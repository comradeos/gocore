package main

import (
	"errors"
	"fmt"
	mp "go000abc/myPackage"
	t "go000abc/tools"
	"math"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var powResult float64 = math.Pow(3, 5)
	fmt.Printf("powres: %.3f\n", powResult)

	var myError error = errors.New("ERR")

	if myError != nil {
		fmt.Println(myError.Error())
	}

	var num int = 2314
	pNum := &num
	fmt.Println(pNum)
	fmt.Println(mp.MyNumber)

	mp.Sep()

	fmt.Println(divZero(12, 2))
	fmt.Println(divZero(12, 0))

	t.Line(5)
	t.Line(15)
	t.Line(25)
	t.Line(0)

	d1 := []byte("hello\ngo\n")
	err := os.WriteFile("dat1", d1, 0644)
	check(err)

	//fmt.Println(5.0 / 2.0)
}

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

/*

go build main.go && ./main

*/
