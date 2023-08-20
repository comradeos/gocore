package main

import (
	"fmt"
	"time"
)

func generateNumbers(n int, ch chan int) {
	for i:=0; i<=n; i++ {
		ch <- i*2
	}
	close(ch) // якщо не закрити канал виникатиме дедлок
}

func main() {
	numbers := make(chan int)

	go generateNumbers(4, numbers)

	for {
		number, ok:= <-numbers

		fmt.Println(number, ok)

		if !ok {
			break
		}

		time.Sleep(time.Second)
	}
}