package chapters

import (
	"crypto/sha256"
	"fmt"
)

func DataTypes2() {
	fmt.Println("DataTypes2")

	a := []byte("Hello")
	b := []byte("Hello")

	sha1 := sha256.Sum256(a)
	sha2 := sha256.Sum256(b)

	fmt.Printf("%x\n", sha1)
	fmt.Printf("%x\n", sha2)
}

func ArraysSlicesMaps() {
	myArray := [3]int{1, 2, 3}
	fmt.Println(myArray)

	mySlice := []int{1, 2, 3}
	fmt.Println(mySlice)
	mySlice = append(mySlice, 4)
	fmt.Println(mySlice)

	myMap := map[string]int{}
	myMap["one"] = 1
	myMap["two"] = 2
	myMap["three"] = 3

	fmt.Println(myMap)

	for key, value := range myMap {
		fmt.Println(key, value)
	}

	fmt.Println("-------------------------")

	type Struct1 struct{}

	var s0 Struct1
	s0 = Struct1{}
	fmt.Println(s0) // {}

	var s1 = Struct1{}
	fmt.Println(s1) // {}

	s2 := Struct1{}
	fmt.Println(s2) // {}

	s3 := new(Struct1)
	fmt.Println(s3) // &{}

	var s4 *Struct1
	fmt.Println(s4) // <nil>
}

func showCap(s *[]int) {
	*s = append(*s, (*s)[len(*s)-1]+1)
	fmt.Println(*s, cap(*s))
}

func LenCap() {
	var mySlice = []int{1}
	fmt.Println(mySlice, cap(mySlice))

	for range 20 {
		showCap(&mySlice)
	}
}
