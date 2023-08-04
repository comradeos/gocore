package main

import "fmt"

var globalVariable string = "I'm the global var!"
const pi = 3.1415

func main() {
	
	var a int = 1; // int, int8/16/32/64, uint, uint8/16/32/64, uintptr, byte
	fmt.Println(a)
	
	var b int8 = 2;
	fmt.Println(b)
	
	localVariable := "I'm local..."
	localVariable = "cool"
	// localVariable := "wrooong"
	fmt.Println(globalVariable)
	fmt.Println(localVariable)

	f64 := float64(5)
	fmt.Printf("f64: %f\n", f64)
} 