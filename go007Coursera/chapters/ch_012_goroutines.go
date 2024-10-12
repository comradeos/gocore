package chapters

import (
	"fmt"
	"runtime"
)

const (
	iterationNum  = 7
	goroutinesNum = 5
)

func Ch012Goroutines1() {
	fmt.Println("============== Hello from Ch012Goroutines() ==============")

	for i := 0; i < goroutinesNum; i++ {
		go doSomeWork(i)
	}

	fmt.Scanln()
}

func doSomeWork(in int) {
	for j := 0; j < iterationNum; j++ {
		fmt.Println(formatWork(in, j))
		runtime.Gosched()
	}
}

func formatWork(in int, j int) string {
	var result = ""

	for i := 0; i < in; i++ {
		result = result + "--|--"
	}

	return result + fmt.Sprintf("> th %d iter %d", in, j)
}
