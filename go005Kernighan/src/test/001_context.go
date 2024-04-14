package test

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func UnderstandingContext() {
	fmt.Println("UnderstandingContext")
	wg := sync.WaitGroup{}
	wg.Add(1)

	ctx, cancel := context.WithCancel(context.Background())

	go BGFunc(&wg, ctx)
	time.Sleep(time.Second * 5)
	cancel()

}

func BGFunc(wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()

	select {
	case <-ctx.Done():
		fmt.Println("BGFunc: Done")
		return
	default:
		fmt.Println("BGFunc: Working")
		for i := range 10 {
			fmt.Println(i)
			time.Sleep(time.Second)
		}
	}
}
