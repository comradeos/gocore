package main

import "fmt"

type A struct {
	Id int
}

type B struct {
	Name string
}

type C struct {
	Timestamp int
}

type D struct {
	Value string
}

type E struct {
	a A
	b B
	c C
	d D
}

func main() {
	var e E
	e.a.Id = 1
	e.b.Name = "Name"
	e.c.Timestamp = 1234567890
	e.d.Value = "Value"
	fmt.Println(e)
}