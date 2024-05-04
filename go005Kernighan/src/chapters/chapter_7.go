package chapters

import "fmt"

type Human interface {
	Speak()
	Run()
}

type Person1 struct {
	Name string
}

func (p *Person1) Speak() {
	fmt.Println("Hello, my name is", p.Name)
}

func (p *Person1) Run() {
	fmt.Println("I'm running")
}

type Mechanism interface {
	Speak()
	Walk()
	Idle()
	Work()
}

type Robot struct {
	Model string
}

func (r *Robot) Speak() {
	fmt.Printf("%s speaking...\n", r.Model)
}

func (r *Robot) Walk() {
	fmt.Printf("%s walking...\n", r.Model)
}

func (r *Robot) Idle() {
	fmt.Printf("%s is idle...\n", r.Model)
}

func (r *Robot) Work() {
	fmt.Printf("%s is working...\n", r.Model)
}

func Interfaces() {
	var h Human = &Person1{"John"}
	h.Speak()
	h.Run()

	var r Mechanism = &Robot{
		"C3PO",
	}

	r.Speak()
	r.Walk()
}
