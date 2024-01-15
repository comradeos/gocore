package modules

import (
	"fmt"
	"sync"
)

func RaceCondition() {
	var x = 0

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := 0; i < 10; i++ {
			x++
			fmt.Printf("goroutine 1: %d\n", x)
		}

		wg.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			x--
			fmt.Printf("goroutine 2: %d\n", x)
		}

		wg.Done()
	}()

	wg.Wait()
	fmt.Println(x)
}
