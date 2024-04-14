package test

import "fmt"

type animal interface {
	makeSound()
}

type Dog struct {
}

func (d *Dog) makeSound() {
	fmt.Println("Bark")
}

func Interfaces() {
	var dog animal = new(Dog)
	dog.makeSound()
}
