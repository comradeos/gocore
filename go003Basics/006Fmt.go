package main

import "fmt"

func main() {
	fmt.Println("Learning FMT:")
	fmt.Print("Just line\n")
	fmt.Printf("digit: %d \n", 25)
	fmt.Printf("just value: %v \n", "pff")
	fmt.Printf("just value: %v \n", 3.31)
	fmt.Printf("just value: %v \n", 2E-4)
	fmt.Printf("just value: %#v \n", 2E-4)
	fmt.Printf("just value: %v \n", "Hello")
	fmt.Printf("just value: %#v \n", "Hello")
	fmt.Println('A') // char 65
}