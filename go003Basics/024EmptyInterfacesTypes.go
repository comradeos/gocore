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

func printType2(value interface{}) {
	switch value.(type) {
		case int:
			f.Println("INT")
		case string:
			f.Println("INT")
		default:
			f.Println("I don't know")

	}
}

func main() {
	f.Println("Empty Interafcaces, type check")
	printType(12)
	printType("hello")
	printType(true)
	
	printType2(12)
	printType2("hello")
	printType2(true)

}