package main

import (
	"go006Bodner/chapters/chapter_2"
	"go006Bodner/chapters/chapter_4"
	"go006Bodner/chapters/chapter_5"
	"go006Bodner/libraries/my_lib_1"
)

func main() {
	my_lib_1.TestLib()
	chapter_2.TypesVars()
	//tests.ST001Overload()
	//tests.ST002NilsEquation()
	chapter_2.CompoundTypes()
	chapter_2.Structs()
	chapter_4.Shadowing()
	chapter_4.OperatorIf()
	chapter_4.ForTypes()
	chapter_4.Labels()
	chapter_4.Switches()
	chapter_4.Goto()
	//chapter_4.ConsoleInput()
	chapter_5.Functions()
}
