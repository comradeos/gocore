package main

import (
	"fmt"
	h "go003Basics/helper"
)

func main() {
	var mySlice = []string {
		"a",
		"b",
		"c",
	}


	fmt.Printf("MySlice length: %d\n", len(mySlice))
	fmt.Printf("MySlice capacity: %d\n", cap(mySlice))
	
	fmt.Println("Let's add some element")

	mySlice = append(mySlice, "d")
	mySlice = append(mySlice, "c")
	mySlice = append(mySlice, "d")
	mySlice = append(mySlice, "f")

	fmt.Printf("New MySlice length: %d\n", len(mySlice))
	fmt.Printf("New MySlice capacity: %d\n", cap(mySlice))
	
	for i,v := range mySlice {
		fmt.Printf("%d: %s\n", i, v)
	}

	h.Sep()
}