package main

import (
	"fmt"
	"time"
)

func someActions() {
	for {
		fmt.Println("Some acts...")
		time.Sleep(time.Second * 1)
	}
}

func main() {
	go func(){
		for i:=1; i<5; i++ {
			fmt.Println("Anon is being doing...")
			time.Sleep(time.Millisecond * 500)
		}
		fmt.Println("Anon goroutine is finished!")
	}()
	someActions()
}