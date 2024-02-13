package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Start")

	var wg sync.WaitGroup

	for i:=0; i<3; i++ {
		wg.Add(1)
		go myFunc(i, &wg)
	}

	wg.Wait()
}

func myFunc(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("myFunc ... %d\n", i)
	time.Sleep(1 * time.Second)
}