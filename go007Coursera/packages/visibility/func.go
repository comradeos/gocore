package visibility

import "fmt"

func CreatePerson(id int, name string) *Person {
	return &Person{
		Id:           id,
		Name:         name,
		privateField: "XXX",
	}
}

func PrintSecret(p *Person) {
	fmt.Println("privateField", p.privateField)
}

func UpdateSecret(p *Person, secret string) {
	p.privateField = secret
}
