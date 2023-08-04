package main

import (
	"fmt"
)

func sep() { fmt.Println("---------------") }

func main() {

	sep()

	var toDoList = [3]string {
		"buy coffee",
		"learn go",
		"play remnant 2",
	}

	for index, value := range toDoList {
		fmt.Printf("%d. %s\n", index, value)
	} 

	sep()

	var toDo2 = [...]string {
		"hello",
		"mf",
		"world",
		"pfff",
	}

	for _, value := range toDo2 {
		fmt.Println(value)
	}
	
	sep()

	for i:=0; i<=10; i++ {
		fmt.Println(i)
	}

	sep()

	var num int
	num = 1
	for {
		num++
		fmt.Println(num)
		if num > 4 { break }
	}

	sep()

	var arr [3]int
	fmt.Println(arr)
}