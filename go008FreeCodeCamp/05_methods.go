package main

import "fmt"

type Rectangle struct {
	Width int
	Height int
}

func (r * Rectangle) area() int {
	return r.Width * r.Height
}

func methods() {
	sep()
	red("05 methods")
	sep()

	rectangle := Rectangle{
		Width: 10,
		Height: 10,
	}

	fmt.Println(rectangle.area())
}

