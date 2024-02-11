package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start")
	data := make(chan int)
	exit := make(chan int)

	go func(){
		for i := 0; i < 10; i++ {
			fmt.Println(i)
			time.Sleep(1 * time.Second)
		}

		exit<-0
	}()

	go selectOne(data, exit)
	
	time.Sleep(20 * time.Second)
}

func selectOne(data, exit chan int) {
	x := 0

	for {
		select {
		case data<-x:
			x++
		case <-exit:
			fmt.Println("Exit")
			return
		}
	}
}