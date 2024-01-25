package main

import (
	"fmt"
)

type Point struct {
	X int
	Y int
}

func (p *Point) info() {
	fmt.Println("X =", p.X, "Y =", p.Y)
}

func main() {
	var my_list = []string {
		"hello",
		"cruel",
		"world",
	} 

	for index, value := range my_list {
		fmt.Println(index, value)
	}

	// игнорируем индекс
	for _, value := range my_list {
		fmt.Println(value)
	}

	// pointMap := map[string] Point { // или make(map[string] Point) // или map[ключ тип]значение тип
	// 	"p1": Point {1, 2},
	// }

	// pointMap := map[string] int {}

	mapPoints := map[string] Point {
		"p1": Point {1, 2},
		"p2": Point {3, 4},
	}

	fmt.Println(mapPoints["p1"].X)
	fmt.Println(mapPoints["p1"].Y)

	inst := mapPoints["p1"]
	inst.info()
}