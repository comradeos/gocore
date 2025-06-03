package main 

import "fmt"

type Vehicle struct {
	Name string
}

type Car struct {
	Vehicle
	Model string
}

func (c Car) PrintModel() {
	fmt.Println("Модель:", c.Model)
}

func main() {
	fmt.Println("inheritance")
	
	var car Car
	car.Name = "Машина"
	car.Model = "Тойота-Хуйота"
	fmt.Printf("Шо: %s, Модель: %s\n", car.Name, car.Model)

	car.PrintModel()
}
