package main

import f "fmt"
import "runtime"

func divByZero(a int, b int) int {
	return a / b
}

func main() {
	runtime.GOMAXPROCS(2)
	cpus := runtime.NumCPU()
	f.Println(cpus)

	defer func(){
		if r := recover(); r != nil {
			f.Printf("шось не так: %s\n", r)
		}
	}()

	result := divByZero(2, 0)
	f.Println(result)
}
