package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		for _, c := range `-\|/` {
			fmt.Printf("\r%c", c)
			time.Sleep(time.Millisecond * 1000)
		} 
	}
}