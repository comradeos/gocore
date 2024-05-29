package tests

import "fmt"

type Printer interface {
	Print()
}

type IntPrinter struct{}

func (ip IntPrinter) Print() {
	fmt.Println("IntPrinter")
}

type StrPrinter struct{}

func (sp StrPrinter) Print() {
	fmt.Println("StrPrinter")
}

func Print(p Printer) {
	p.Print()
}

func ST001Overload() {
	var p1 IntPrinter = IntPrinter{}
	var p2 StrPrinter = StrPrinter{}
	Print(p1)
	Print(p2)
}
