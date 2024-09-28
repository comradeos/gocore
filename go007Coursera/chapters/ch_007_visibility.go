package chapters

import (
	"fmt"
	"go007Coursera/packages/visibility"
)

func Ch007Visibility() {
	fmt.Println("============== Hello from Ch007Visibility() ==============")

	var p = visibility.CreatePerson(1, "John")
	fmt.Println("PublicVar", visibility.PublicVar)
	visibility.PrintSecret(p)

	var p2 visibility.Person = visibility.Person{
		Id:   2,
		Name: "Jane",
	}
	visibility.PrintSecret(&p2)
}
