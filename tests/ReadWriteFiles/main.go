package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("ReadWriteFiles Starting...")

	// Чтение файла
	data, err := ioutil.ReadFile("./test.txt")
	if err != nil { fmt.Println(err) }
	fmt.Println(data)
	

	// Запись в файл
	err = ioutil.WriteFile("./test2.txt", data, 0644)
}