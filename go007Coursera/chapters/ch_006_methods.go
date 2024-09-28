package chapters

import "fmt"

func Ch006Methods() {
	fmt.Println("============== Hello from Ch006Methods() ==============")

	// методы

	human1 := Human{1, "John"}
	human1.SayHello()

	human1.SetName1("Alice")
	fmt.Println("human1", human1) // не изменилось, так как SetName1 принимает копию

	human1.SetName2("Bob")
	fmt.Println("human1", human1) // изменилось, так как SetName2 принимает по ссылке

	fmt.Println("-----------------")

	fmt.Printf("%#v\n", human1) // %#v - выводит структуру
}

type Human struct {
	Id   int
	Name string
}

func (h *Human) SayHello() { // перелача по ссылке
	fmt.Println("Hello, my name is", h.Name)
}

func (h Human) SetName1(name string) {
	h.Name = name
}

func (h *Human) SetName2(name string) {
	h.Name = name
}
