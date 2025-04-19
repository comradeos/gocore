package main

import "fmt"

type IShape interface {
	area() float64
	perimetr() float64
}

type Square struct {
	Side float64
}

func (s *Square) area() float64 {
	return s.Side * s.Side
}

func (s *Square) perimetr() float64 {
	return 4 * s.Side
}


func interfaces() {
	sep()
	red("06 interfaces")
	sep()

	var square IShape = &Square{
		Side: 5,
	}

	fmt.Println(square.area())
}

