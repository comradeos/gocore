package modules

import "fmt"

type Animal struct {
	Name string
}
type Cat struct {
	Animal
}
type Dog struct {
	Animal
}

func (a Animal) Sound() {
	fmt.Printf("Animal sound...\n")
}

func (c Cat) Sound() {
	fmt.Printf("%s makes cat sound...\n", c.Name)
}

func (d Dog) Sound() {
	fmt.Printf("%s makes dog sound...\n", d.Name)
}

func GoPolymorphism() {
	cat := Cat{Animal{Name: "Kotski"}}
	cat.Sound()

	dog := Dog{Animal{Name: "Sobaken"}}
	dog.Sound()
}
