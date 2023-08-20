package main

import (
	"fmt"
	"time"
)

func setProgramLifetime(seconds int) {
	for i:=0; i<seconds; i++ {
		fmt.Println("...")
		time.Sleep(time.Second * 1)
	}
	fmt.Println("Program finished!")
}

func calculateSomthing(res chan int) {
	result := 0
	for i:=0; i<3; i++ {
		fmt.Println("calc1 doing...")
		result += 1
		time.Sleep(time.Second * 1)
	}
	res<-result
}

func main() {
	chanResult := make(chan int)

	go calculateSomthing(chanResult)

	fmt.Println("a")
	fmt.Println(<-chanResult)
	setProgramLifetime(5)
}