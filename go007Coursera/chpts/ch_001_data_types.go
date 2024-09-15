package chpts

import (
	"fmt"
	"unicode/utf8"
)

func Ch001DataTypes() {
	fmt.Println("Hello from Ch001DataTypes()")

	var num0 int // 0 by default
	var num1 int = 1
	var num2 = 20
	num3 := 30

	fmt.Println(num0, num1, num2, num3)

	var numIndex = 25
	//++numIndex // Go doesn't support ++ or -- operators

	numIndex++
	fmt.Println("numIndex", numIndex)

	numIndex--
	fmt.Println("numIndex", numIndex)

	weight, height := 70, 180
	fmt.Println("weight", weight, "height", height)

	weight, age := 80, 30
	fmt.Println("weight", weight, "age", age)

	// int платформозависимый тип, который может быть 32 или 64 бита в зависимости от архитектуры
	var i int = 1 // плаформа выбрана автоматически
	fmt.Println("i", i)

	// int8, int16, int32, int64 - максимальные значения
	var i8 int8 = 127
	var i16 int16 = 32767
	var i32 int32 = 2147483647
	var i64 int64 = 9223372036854775807
	fmt.Println("i8", i8, "i16", i16, "i32", i32, "i64", i64)

	i64++
	fmt.Println("i64", i64) // overflow

	// uint8, uint16, uint32, uint64 - максимальные значения
	var ui8 uint8 = 255
	var ui16 uint16 = 65535
	var ui32 uint32 = 4294967295
	var ui64 uint64 = 18446744073709551615
	fmt.Println("ui8", ui8, "ui16", ui16, "ui32", ui32, "ui64", ui64)

	// float32, float64
	var f32 float32 = 3.14
	var f64 float64 = 3.14
	fmt.Println("f32", f32, "f64", f64)

	// bool
	var b bool = true
	fmt.Println("b", b)

	// complex64, complex128
	var c64 complex64 = 1 + 2i
	var c128 complex128 = 1 + 2i
	fmt.Println("c64", c64, "c128", c128)

	// string
	var s1 string = "Hello, World!"
	fmt.Println("s1", s1)

	var s2 string = `Hello, World!\n`
	fmt.Println("s2", s2)

	var s3 string = `Hello,
	World!`
	fmt.Println("s3", s3)

	// byte (uint8)
	var by1 byte = 'A'
	fmt.Println("by1", by1)

	var by2 byte = '\x41'
	fmt.Println("by2", by2)

	// rune (int32)
	var r1 rune = 'A'
	fmt.Println("r1", r1)

	var r2 rune = '馬'
	fmt.Println("r2", r2)

	// вывести как строку
	fmt.Println(string(r2))

	// длина строки в байтах
	var helloWorld = "Hello, World!"
	fmt.Println("len(helloWorld)", len(helloWorld))

	var helloWorldJapanese = "こんにちは世界！"                             // 8 байт?
	fmt.Println("len(helloWorldJapanese)", len(helloWorldJapanese)) // 24 байта, потому что UTF-8, один символ может занимать несколько байт

	var helloWorldJapaneseRune = []rune("こんにちは世界！")                         // 8 символов
	fmt.Println("len(helloWorldJapaneseRune)", len(helloWorldJapaneseRune)) // 8

	// RunCountInString
	var c1 = utf8.RuneCountInString(helloWorldJapanese)
	fmt.Println("utf8.RuneCountInString(helloWorldJapanese) c1", c1) // 8

	// конвертация в slice байт и обратно
	var helloWorldBytes = []byte(helloWorld)
	fmt.Println("helloWorldBytes", helloWorldBytes) // [72 101 108 108 111 44 32 87 111 114 108 100 33]

	var helloWorldString = string(helloWorldBytes)
	fmt.Println("helloWorldString", helloWorldString) // Hello, World!

	// константы

	const pi = 3.14
	fmt.Println("pi", pi)

	// блок констант

	const (
		pi1 = 3.14159265359
		pi2 = 3.14159265359
	)
	fmt.Println("pi1", pi1, "pi2", pi2)

	// iota

	const (
		zero = iota
		one
		two
		_
		_
		five
	)

	fmt.Println("zero", zero, "one", one, "two", two, "five", five)

	// нетипизированные константы
	const (
		year         = 2024
		yearType int = 2025
	)

	fmt.Println("year", year, "yearType", yearType)
	fmt.Println("year + yearType", year+yearType)

	var strToAdd = "Hello, " + "World!"
	fmt.Println("strToAdd", strToAdd)
	//fmt.Println("strToAdd + yearTyped", strToAdd+yearType) // invalid operation: strToAdd + yearType (mismatched types string and int)

	// приведение типов

	var idx int = 1

	type UserId int

	var userId UserId = UserId(idx)
	fmt.Println("userId", userId)

	// указатели

	var a = 7     // переменная типа int
	var aPtr = &a // указатель на переменную a
	*aPtr = 8     // разыменование указателя и присвоение значения

	fmt.Println("a", a) // 8

}
