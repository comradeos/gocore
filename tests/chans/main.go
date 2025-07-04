package main 

import (
	"fmt"
	"time"
)

func myGor(ch chan string) {
	counter := 0
	for {
		counter++
		ch<-fmt.Sprintf("msg %d\n", counter)
		time.Sleep(1*time.Second)
	}
}

func main() {
	println("chans")

	ch := make(chan string)

	go myGor(ch)

	for {
		time.Sleep(1*time.Second)
		println(<-ch)
	}
}