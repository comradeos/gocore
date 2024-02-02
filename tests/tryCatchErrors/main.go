package main

import (
    "fmt"
    "os"
    "path/filepath"
)

func main() {
    // Получаем путь к выполняемому файлу
    exe, err := os.Executable()
    if err != nil {
        // Если произошла ошибка, выводим ее и завершаем программу
        fmt.Println("Error:", err)
        return
    }

    // Используем filepath.Dir для получения каталога, в котором находится файл
    dir := filepath.Dir(exe)

    // Изменяем текущий рабочий каталог на каталог выполняемого файла
    err = os.Chdir(dir)
    if err != nil {
        // Если произошла ошибка, выводим ее и завершаем программу
        fmt.Println("Error:", err)
        return
    }

    // Выводим новый рабочий каталог
    fmt.Println("Changed working directory to:", dir)

    dir, err = os.Getwd()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

	fmt.Println(dir)

	_, err = os.Open("non-existing-file.txt")
    if err != nil {
        // Обработка ошибки
        fmt.Println("An error occurred:", err)
        return
    }

    // Если ошибки нет, продолжаем выполнение программы
    fmt.Println("No error occurred")
}