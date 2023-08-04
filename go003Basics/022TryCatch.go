package main

import f "fmt"

func div(a int, b int) int {
	defer func(i int) {
		if r := recover();
		r != nil {
			f.Println("ERR!")
			f.Println(i)
		}
	}(777)
	return a / b
}

func main() {
	f.Println(div(1,0))
}
