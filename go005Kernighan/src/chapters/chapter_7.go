package chapters

import (
	"fmt"
	"sort"
)

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

type Model struct {
	Name string
	Id   int
}

type ByName []*Model

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

func Sorting() {
	var m1 *Model = &Model{Name: "Abc1", Id: 9000}
	var m2 *Model = &Model{Name: "Bcd2", Id: 8000}
	var m3 *Model = &Model{Name: "Cde3", Id: 7000}

	var list []*Model

	list = append(list, m3)
	list = append(list, m2)
	list = append(list, m1)

	//sort.Sort(ByName(list))

	sort.Slice(list, func(i, j int) bool {
		return list[i].Name < list[j].Name
	})

	fmt.Println(list[0].Name)
	fmt.Println(list[1].Name)
	fmt.Println(list[2].Name)

}
