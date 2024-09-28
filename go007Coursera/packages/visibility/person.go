package visibility

import "fmt"

var PublicVar = "Public var"
var private = "Private var"

type Person struct {
	Id           int
	Name         string
	privateField string
}

func (p *Person) SayHello() {
	fmt.Println("Hello, my name is", p.Name)
}
