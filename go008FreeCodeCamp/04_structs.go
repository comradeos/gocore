package main

import "fmt"

type Car struct {
	Brand string
	Model string
	Year int	
}

func structs() {
	sep()
	red("04 structs")
	sep()

	car := Car{}
	car.Brand = "BWM"
	car.Model = "M7"

	fmt.Println(car.Model)
}

