package main

import (
	"fmt"
)

func main() {
	fmt.Println("Interfaces")

	var mazda ICar 
	
	mazda = &Mazda {
		Year: 2020,
		Model: "CX-5",
		Color: "Red",
	}
	
	var info string = mazda.Info()
	
	fmt.Println(info)

	mazda.Drive()
}

type ICar interface {
	Info() string
	Drive()
}

type Mazda struct {
	Year int
	Model string
	Color string
}

func (mazda * Mazda) Info() string {
	return fmt.Sprintf("Mazda %s, %s, %d", mazda.Model, mazda.Color, mazda.Year)
}

func (mazda * Mazda) Drive() {
	fmt.Println("Mazda is driving")
}