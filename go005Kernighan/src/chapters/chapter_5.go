package chapters

import "fmt"

type Speaker interface {
	Speak()
	Run()
}

type Person struct {
	Name string
}

func (p *Person) Speak() {
	fmt.Println("Hello, my name is", p.Name)
}

func (p *Person) Run() {
	fmt.Println("I'm running")
}

func MyRec(num int) {
	num += num
	fmt.Println(num)
	if num > 100 {
		return
	}
	MyRec(num)
}

func Functions() {
	var s Speaker
	s = &Person{"John"}
	s.Speak()
	s.Run()
	MyRec(1)
}
