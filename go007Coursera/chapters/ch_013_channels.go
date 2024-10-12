package chapters

import (
	"fmt"
)

func Ch013Channels() {
	ch1 := make(chan int, 1)

	fmt.Println("============== Hello from Ch013Channels() ==============")

	go func(in chan int) {
		value := <-in
		fmt.Println("GO get from chan", value)
		fmt.Println("GO after read from chan")
	}(ch1)

	ch1 <- 222
	ch1 <- 111

	fmt.Println("GO MAIN after put in chan")

	fmt.Scanln()
}
