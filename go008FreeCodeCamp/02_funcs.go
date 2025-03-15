package main

import "fmt"

func functions() {
	sep()
	red("02 functions")
	sep()

	myStr := concat("Hello", "World")

	fmt.Println(myStr)
}

func concat(a string, b string) string {
	return a + b
}
