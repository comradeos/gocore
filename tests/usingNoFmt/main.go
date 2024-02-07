package main

func main() {
	println(123)
	println("asd")

	person := Person{Id: 1, Name: "Ivan"} 

	println(person.Id)
	println(person.Name)

}

type Person struct {
	Id int
	Name string
}
