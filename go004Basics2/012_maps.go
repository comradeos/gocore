package main

import (
	"fmt"
)

type Key struct {
	NAME string
	VALUE int
}

type Value struct {
	NAME string
	VALUE int
}


type Point1 struct {
	X int
	Y int
}

type Point2 struct {
	x int
	y int
}

func main() {
	fmt.Println("Maps")

	myMap := map[Key] Value {}

	key1 := Key {"key1", 1}

	value1 := Value {"value1", 1}

	myMap[key1] = value1

	fmt.Println(myMap[key1].NAME)

	for key, value := range myMap {
		fmt.Println(key.NAME, value.NAME)
	}
}