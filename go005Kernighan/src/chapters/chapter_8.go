package chapters

import (
	"fmt"
	"sync"
	"time"
)

func SomeAct1(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range 10 {
		time.Sleep(time.Second / 3)
		fmt.Printf("SomeAct1: %d\n", i)
	}
}

func SomeAct2(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range 10 {
		time.Sleep(time.Second / 2)
		fmt.Printf("SomeAct2: %d\n", i)
	}
}

func Goroutines() {
	fmt.Println("Goroutines")

	var wg sync.WaitGroup

	wg.Add(2)
	go SomeAct1(&wg)
	go SomeAct2(&wg)

	wg.Wait()
}

// ------------------------------------------------------------

func ChanSend(wg *sync.WaitGroup, ch chan MyChannel) {
	defer wg.Done()
	mc := MyChannel{Value: "Hello", Data: 0}
	for i := range 10 {
		time.Sleep(time.Second * 2)
		mc.Data = i
		ch <- mc // send i to channel ch
	}
}

func ChanReceive(wg *sync.WaitGroup, ch chan MyChannel) {
	defer wg.Done()
	for _ = range 10 {
		fmt.Printf("ChanReceive: %d\n", (<-ch).Data) // receive from ch
	}
}

type MyChannel struct {
	Value string
	Data  int
}

func Channels() {
	fmt.Println("Channels")

	var wg sync.WaitGroup
	wg.Add(2)

	ch := make(chan MyChannel)

	go ChanSend(&wg, ch)
	go ChanReceive(&wg, ch)

	wg.Wait()

}
