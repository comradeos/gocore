package main

import "fmt"

func main() {
	data := []byte{0xAA, 0x03, 0xFF}
	checksum := GetCheckSum(data)
	prettyPrint(checksum)
}

func GetCheckSum(data []byte) byte {
	var sum byte

	fmt.Printf("SUM: %d\n", sum)

    for _, b := range data[1:] {
		fmt.Printf("%d + %d = ", sum, b)
		sum += b
		fmt.Printf("%d\n", sum)
    }

	result := byte(0 - sum)

	fmt.Printf("0 - %d = %d\n", sum, result)

    return result
}

func prettyPrint(b byte) {
	fmt.Printf("\nDecimal: %d | Hexadecimal: 0x%02X\n\n", b, b)
}

// Модульная арифметика: При вычитании 0 - 172 получаем -172. 
// Чтобы привести этот результат к допустимому диапазону byte, добавляем 256:
// 0 − 172 + 256 = 84