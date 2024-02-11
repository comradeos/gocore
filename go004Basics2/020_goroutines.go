package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start")
	ch1 := make(chan int)
	go goTest(ch1)

	for i := range ch1 {
		fmt.Println(i)
	}
}

func goTest(ch1 chan int) {
	for i := 0; i < 10; i++ {
		ch1 <- i
		time.Sleep(1 * time.Second)
	}
	close(ch1)
}