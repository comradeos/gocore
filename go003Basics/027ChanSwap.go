package main

import (
	"fmt"
	"time"
)

func goroutine1(gr1 chan string) {
	for i:=0; i<3; i++ {
		fmt.Println("goroutine1...")
		time.Sleep(time.Second)
	}
	gr1<-"goroutine1RESULT"
}

func goroutine2(gr1 chan string, gr2 chan string) {
	for i:=0; i<5; i++ {
		fmt.Println("...goroutine2")
		time.Sleep(time.Second)
	}
	gr2<-<-gr1
}

func main() {
	ch1 := make(chan string)
	go goroutine1(ch1)

	ch2 := make(chan string)
	go goroutine2(ch1, ch2)
	
	fmt.Println(<-ch2)

	time.Sleep(time.Second * 7)
}