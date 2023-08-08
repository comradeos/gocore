package main

import f "fmt"

func divByZero(a int, b int) int {
	return a / b
}

func main() {
	defer func(){
		if r := recover(); r != nil {
			f.Printf("шось не так: %s\n", r)
		}
	}()

	result := divByZero(2, 0)
	f.Println(result)
}
