package main

import (
	"MyProject/MyPackages/PackageA"
	"fmt"
	"github.com/pkg/errors"
)

func main() {

	fmt.Println("Hello, world!")

	PackageA.FunctionA()

	err := errors.New("this is an error")
	
    fmt.Println(err)
}
