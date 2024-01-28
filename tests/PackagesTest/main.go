// main.go
package main

import (
    "myPackage"
)

func main() {
    myPackage.MyFunction()
	myPackage.AnotherFunction()
	myStruct := myPackage.GetStruct()
	println(myStruct.MyField)
}