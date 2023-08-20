package main

import (
    "fmt"
    "math/big"
    "time"
)

func waitSpinner(result chan big.Int) {
	for {
		for _, c := range `-\|/` {
			fmt.Printf("\r%c", c)
			time.Sleep(time.Millisecond * 1000)
		}
	}
}

func factorial(x *big.Int) *big.Int {
    n := big.NewInt(1)
    if x.Cmp(big.NewInt(0)) == 0 {
        return n
    }
    return n.Mul(x, factorial(n.Sub(x, n)))
}

func main() {
	var res = make(chan big.Int)
	
	if *res == nil {
		fmt.Println(123)
	}

    fmt.Println(factorial(big.NewInt(45)))
}