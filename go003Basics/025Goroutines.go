package main

import (
	"fmt"
	"time"
)

func someAction(name string, seconds int) {
	timeStart := time.Now()

	for i:=1; i<=seconds; i++ {
		fmt.Printf("Doing %s ...\n", name)
		time.Sleep(time.Second * 1)
	}
	
	timeSpent := time.Since(timeStart)
	fmt.Printf("Action %s is done in %v!\n", name, timeSpent)
}

func main() {
	timeStart := time.Now()

	go someAction("act1", 2) // run in background

	go func() { // анонимная горутина
		for i:=1; i<=5; i++ {
			fmt.Println("Anonym gouroutine...")
			time.Sleep(time.Second)
		}
		fmt.Println("Anon is done!")
	}()
	
	someAction("act2", 5)

	timeSpent := time.Since(timeStart)
	fmt.Println(timeSpent)
}