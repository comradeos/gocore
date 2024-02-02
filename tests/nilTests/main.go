package main

import (
	"fmt"
)

func main() {
	var a *int;
	var b *string;

	fmt.Println(a == nil);
	fmt.Println(b == nil);
	fmt.Println(a == b); // ошибка 
}