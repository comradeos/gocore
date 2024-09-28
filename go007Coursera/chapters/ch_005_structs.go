package chapters

import "fmt"

func Ch005Structs() {
	fmt.Println("============== Hello from Ch005Structs() ==============")

	// структуры

	var person1 Person
	person1.Id = 1
	person1.Name = "John"

	fmt.Println("person1", person1)

	var person2 = Person{Id: 2, Name: "Alice"}
	fmt.Println("person2", person2)

	person3 := Person{3, "Bob"}
	fmt.Println("person3", person3)

	var pPerson4 *Person
	fmt.Println("person4", pPerson4)

	var account1 Account
	account1.Id = 1
	account1.Cleaner = func(str string) string {
		return "cleaned: " + str
	}
	account1.Owner = person1
	fmt.Println("account1", account1)

	res := account1.Cleaner("some string")
	fmt.Println("res", res)

	// вложенные структуры

	var structB1 StructB
	structB1.FieldA1 = "FieldA1"
	structB1.FieldB1 = "FieldB1"

	fmt.Println("structB1", structB1)

	var structB2 = StructB{
		StructA: StructA{
			FieldA1: "FieldA1",
		},
		FieldB1: "FieldB1",
	}

	fmt.Println("structB2", structB2)

	// рекурсивные структуры

	var structC1 StructC
	structC1.FieldC1 = "FieldC1"
	structC1.StructC = &structC1

	fmt.Println("structC1", structC1.StructC.StructC.StructC.StructC.StructC.StructC.StructC.StructC.FieldC1)

	// анонимные структуры

	var structD = struct {
		FieldD1 string
		FieldD2 int
	}{} // создание экземпляра анонимной структуры

	structD.FieldD1 = "FieldD1"
	structD.FieldD2 = 2
	fmt.Println("structD", structD)

}

type Person struct {
	Id   int
	Name string
}

type Account struct {
	Id      int
	Cleaner func(string) string
	Owner   Person
}

type StructA struct {
	FieldA1 string
}

type StructB struct {
	StructA
	FieldB1 string
}

type StructC struct {
	*StructC
	FieldC1 string
}
