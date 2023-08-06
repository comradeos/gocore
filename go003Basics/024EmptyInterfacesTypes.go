package main

import f "fmt"

func printType(value interface{}) {
	if _, ok := value.(int); ok { 
		f.Println("это число") 
	} 
	if _, ok := value.(string); ok {
		f.Println("это строка")
	} else {
		f.Println("хз что за тип")
	}
}

func main() {
	f.Println("Empty Interafcaces, type check")
	printType(12)
	printType("hello")
	printType(true)
}