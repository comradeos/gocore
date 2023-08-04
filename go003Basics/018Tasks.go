package main

import f"fmt"
import h"go003Basics/helper"

func main() {
	f.Println("Custom types:")
	h.Sep()

	var myMap map[string]int
	myMap = make(map[string]int, 0)
	myMap["a1"] = 123
	myMap["a2"] = 234
	myMap["a3"] = 345
	f.Println(myMap == nil)
	f.Println(myMap["a1"])
	f.Println(len(myMap))

	var myMap2 map[string]int
	f.Println(myMap2 == nil)

}