package main

import (
    "errors"
    "fmt"
)

func main() {
    result, err := div(1, 0)
    if err != nil {
        // Если произошла ошибка, выводим ее
        fmt.Println("Error:", err)
    } else {
        // Если ошибки нет, выводим результат
        fmt.Println("Result:", *result)
    }
}

func div(a int, b int) (*int, error) {
    // Если делитель равен нулю, возвращаем ошибку
    if b == 0 {
        return nil, errors.New("division by zero")
    }

    // Если делитель не равен нулю, выполняем деление
    result := a / b
    return &result, nil
}