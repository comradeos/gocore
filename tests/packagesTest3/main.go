package main

import (
	"fmt"
	"PackageTest3/PackageA"
)

func main() {
	fmt.Println(123)
	PackageA.HelloFromPackageA()
	// PackageA.thisIsPrivate() // ./main.go:11:11: undefined: PackageA.thisIsPrivate
}