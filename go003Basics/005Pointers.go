package main

import (
	"fmt"
)

func main() {
	var num int = 234
	pNum1 := &num
	var pNum2 *int = &num

	fmt.Println(num)
	fmt.Println(&num)
	fmt.Println(pNum1)
	fmt.Println(pNum2)
	
	fmt.Println("-----------")
	fmt.Println(&pNum1)
	fmt.Println(&pNum2)
	
	fmt.Println("-----------")
	*pNum1+=10
	fmt.Println(num)
}