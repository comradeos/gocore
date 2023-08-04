package main

import (
	"fmt"
	"errors"
	"math"
	"os"
)

const PI = math.Pi

func main() {
	fmt.Println("Tasks:")
	var err error = errors.New("I'm an error!")
	fmt.Println(err.Error())
	fmt.Errorf("fmt.Error!\n")

	pNum := new(int)
	*pNum = 123
	ppNum := &pNum
	fmt.Println(pNum)
	fmt.Println(*pNum)
	fmt.Println(**ppNum)


	os.Exit(0)
}

