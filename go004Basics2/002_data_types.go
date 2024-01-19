package main

import (
	"fmt"
)

func main() {
	var a int8 = 127 // -128 to 127 1 byte
	fmt.Printf("%d\n", a)

	var b int16 = 32767 // -32768 to 32767 2 bytes
	fmt.Printf("%d\n", b)

	var c int32 = 2147483647 // -2147483648 to 2147483647 4 bytes
	fmt.Printf("%d\n", c)

	var d int64 = 9223372036854775807 // -9223372036854775808 to 9223372036854775807 8 bytes
	fmt.Printf("%d\n", d)

	var e uint8 = 255 // 0 to 255 1 byte
	fmt.Printf("%d\n", e)

	var f uint16 = 65535 // 0 to 65535 2 bytes
	fmt.Printf("%d\n", f)

	var g uint32 = 4294967295 // 0 to 4294967295 4 bytes
	fmt.Printf("%d\n", g)

	var h uint64 = 18446744073709551615 // 0 to 18446744073709551615 8 bytes
	fmt.Printf("%d\n", h)

	var i int = 9223372036854775807 // -9223372036854775808 to 9223372036854775807 8 bytes or 4 bytes depending on the system (alias for int32 or int64)
	fmt.Printf("%d\n", i)

	var j byte = 255 // 0 to 255 1 byte (alias for uint8)
	fmt.Printf("%d\n", j)

	var k rune = 2147483647 // -2147483648 to 2147483647 4 bytes (alias for int32)
	fmt.Printf("%d\n", k)

	var l uint = 18446744073709551615 // 0 to 18446744073709551615 8 bytes or 4 bytes depending on the system (alias for uint32 or uint64)
	fmt.Printf("%d\n", l)

	var m float32 = 3.40282346638528859811704183484516925440e+38 // 1.401298464324817070923729583289916131280e-45 to 3.40282346638528859811704183484516925440e+38 4 bytes
	fmt.Printf("%f\n", m)

	var n float64 = 1.797693134862315708145274237317043567981e+308 // 4.940656458412465441765687928682213723651e-324 to 1.797693134862315708145274237317043567981e+308 8 bytes
	fmt.Printf("%f\n", n)

	var o complex64 = 1+2i // float32 real and imaginary parts 8 bytes
	fmt.Printf("%f\n", o)

	var p complex128 = 2+3i // float64 real and imaginary parts 16 bytes
	fmt.Printf("%f\n", p)

	var q bool = true // true or false 1 byte
	fmt.Printf("%t\n", q)

	var r string = "Hello World" // string of UTF-8 characters
	fmt.Printf("%s\n", r)

	var s [5]int = [5]int{1, 2, 3, 4, 5} // array of 5 integers
	fmt.Printf("%d\n", s)

	var t []int = []int{1, 2, 3, 4, 5, 123} // slice of integers
	fmt.Printf("%d\n", t)


	var name string = "John Doe" // type inference
	fmt.Printf("%s\n", name)

	name2 := "John Doe" // type inference
	fmt.Printf("%s\n", name2)

	fmt.Print("--------------------\n")

	var newName, age = "John Doe", 30 // type inference
	var result string = fmt.Sprintf("%s %d\n", newName, age) // формирование строки с любым типом данных
	fmt.Printf("%s", result)
}