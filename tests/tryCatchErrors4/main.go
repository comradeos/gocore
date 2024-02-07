package main

import (
    "fmt"
)

func div(a int, b int) int {
    return a / b
}

func main() {
    var res1, res2 int

    func() {
        defer func() {
            if r := recover(); r != nil {
                res1 = 5
            }
        }()
        res1 = div(1, 0)
    }()

    func() {
        defer func() {
            if r := recover(); r != nil {
                res2 = 5
            }
        }()
        res2 = div(1,0)
    }()

    res3 := res1 + res2

    fmt.Println(res3)
}