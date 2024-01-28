package myPackage

import "fmt"

type MyStruct struct {
    MyField string
}

func AnotherFunction() {
    fmt.Println("Hello from myPackage -> AnotherFunction!")
}

func GetStruct() MyStruct {
    return MyStruct{MyField: "This is MyStruct -> Field!"}
}