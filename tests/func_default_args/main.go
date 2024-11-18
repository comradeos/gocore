package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("func_default_args")

	var o Opts
	fmt.Println("o:", o)

	TestFuncDefaultArgs(&o)
}

func TestFuncDefaultArgs(o *Opts) {
	if o == nil || o.Delay == 0 {
		o.Delay = 1
	}

	fmt.Println("A")
	
	time.Sleep(time.Second * 2)

	fmt.Println("B")
}

type Opts struct {
	Delay int
}