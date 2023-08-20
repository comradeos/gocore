package main

import (
	"fmt"
	"time"
)

func genNums(count int, res chan int) {
	for i:=0; i<=count; i++ {
		res <- i*2
	}
	close(res)
}

func main() {
	numbers := make(chan int)

	go genNums(3, numbers)
	
	select {
		case number := <-numbers:
			fmt.Println(number)
		default:
			fmt.Println("no values: channel is closed")
	}

	time.Sleep(time.Second * 100)
}