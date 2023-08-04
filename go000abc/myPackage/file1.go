package myPackage

import "fmt"

const MyNumber int = 275
const myString string = "Hello!"

// go: link
func getMyNumber() int {
	return MyNumber
}

func Sep() {
	fmt.Println("---------------------")
}
