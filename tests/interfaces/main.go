package main

import "fmt"

type Human interface {
	SayHello()
}

type DeutschlandMan struct {
	Name string
}

func (d DeutschlandMan) SayHello() {
	fmt.Printf("gutentak! %s\n", d.Name)
}

type JapanMan struct {
	Name string
}

func (j JapanMan) SayHello() {
	fmt.Printf("konichiwa! %s\n", j.Name)
}

func SayHello(h Human) {
	h.SayHello()
}

func main() {
	var human Human

	human = DeutschlandMan{"DeutschlandMan"}
	SayHello(human)
}


