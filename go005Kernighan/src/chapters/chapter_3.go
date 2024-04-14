package chapters

import (
	"fmt"
	"strings"
	"time"
	"unsafe"
)

func DataTypes() {
	fmt.Println("DataTypes")
	// фундаментальные типы данных (числа, строки и булевы значения)
	var n int = 10
	fmt.Println(unsafe.Sizeof(n) * 8)

	fmt.Printf("%b\n", 7)
	fmt.Printf("%b\n", -7)

	var n2 int8 = 8
	fmt.Printf("bin: %b, dec: %d\n", n2, n2)
	n2 = n2 >> 1
	fmt.Printf("bin: %b, dec: %d\n", n2, n2)

	// составные типы данных (массивы, структуры)
	// ссылочные типы данных (указатели, срезы, карты, каналы, функции)
	// типы интерфейсов
}

func CarriageReturn() {
	for i := 0; i < 10; i++ {
		fmt.Printf("\r%s", strings.Repeat("=", i))
		time.Sleep(time.Millisecond * 100)
	}
}
