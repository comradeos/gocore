package main

import "fmt"

func sep() {
	fmt.Println("--------------------")
}

func red(str string) {
	fmt.Println("\033[31m", str, "\033[0m")
}

func green(str string) {
	fmt.Println("\033[32m", str, "\033[0m")
}

func blue(str string) {
	fmt.Println("\033[34m", str, "\033[0m")
}
