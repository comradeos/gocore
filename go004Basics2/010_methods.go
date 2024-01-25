package main

import (
	"fmt"
)

type Point struct {
	X int
	Y int
}

func movePoint (point Point, x int, y int) Point {
	point.X = x
	point.Y = y
	return point
}

func (point *Point) movePoint (x int, y int) {
	point.X = x
	point.Y = y
}

func main() {
	var p1 Point = Point {1, 2}
	p1 = movePoint(p1, 10, 10)

	p1.movePoint(20, 20)
	fmt.Println(p1.X)
}