package chapters

import (
	"fmt"
)

func Ch002DataTypes2() {
	fmt.Println("============== Hello from Ch001DataTypes2() ==============")

	// массив

	var arr1 [5]int // [0 0 0 0 0]
	var arr2 [5]int = [5]int{1, 2, 3, 4, 5}

	fmt.Println("arr1", arr1)
	fmt.Println("arr2", arr2)

	const size = 5
	var arr3 = [size]int{21, 22, 23, 24, 2 + 5}
	fmt.Println("arr3", arr3)

	// опредилить размер массива по количеству элементов
	arr4 := [...]int{31, 32, 33, 34, 35, 36, 37, 38, 39, 40}
	fmt.Println("arr4", arr4)

	// срез

	var slice1 []int = []int{1, 2, 3, 4, 5}

	fmt.Println("slice1", slice1)

	// append

	slice1 = append(slice1, 6)
	fmt.Println("slice1", slice1)

	var slice2 = make([]int, 5, 10)
	fmt.Println("slice2", slice2)

	var slice3 = make([]int, 15)
	fmt.Println("slice3", slice3)

	slice2 = append(slice2, slice3...)
	fmt.Println("slice2", slice2)

	var sliceA = []int{1, 2, 3, 4, 5}
	fmt.Println("sliceA", sliceA)
	var sliceB = sliceA[:]
	fmt.Println("sliceB", sliceB)
	sliceB[0] = 777
	fmt.Println("sliceA", sliceA)
	fmt.Println("sliceB", sliceB)

	sliceB = append(sliceB, 6)
	sliceB[0] = 888
	fmt.Println("sliceA", sliceA)
	fmt.Println("sliceB", sliceB)

	// копирование среза
	var sliceToCopy = []int{1, 2, 3, 4, 5}
	var sliceCopy = make([]int, len(sliceToCopy), cap(sliceToCopy))
	copy(sliceCopy, sliceToCopy)
	fmt.Println("sliceToCopy", sliceToCopy)
	fmt.Println("sliceCopy", sliceCopy)

	var ints = []int{6, 6, 6, 6, 6, 6}
	fmt.Println("ints", ints)

	copy(ints[1:5], []int{7, 7, 7, 7})
	fmt.Println("ints", ints)

	// map

	var map1 = map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}

	fmt.Println("map1", map1)

	var map2 = make(map[int]map[int]string)
	map2[1] = map[int]string{
		1: "one",
		2: "two",
	}
	map2[2] = map[int]string{
		1: "one",
		2: "two",
	}
	fmt.Println("map2", map2)

	var map3 = make(map[string]interface{})
	map3["one"] = 1
	map3["two"] = "two"
	map3["three"] = []int{1, 2, 3}

	fmt.Println("map3", map3)
	fmt.Println(`map3["one"]`, map3["one"])
	fmt.Println(`map3["two"]`, map3["two"])
	fmt.Println(`map3["three"]`, map3["three"])
	fmt.Println(`map3["four"]`, map3["four"])

	// прзнак наличия ключа в map
	if _, ok := map3["four"]; !ok {
		fmt.Println("map3[\"four\"] not found")
	}

	// удаление ключа из map

	delete(map3, "one")
	fmt.Println("map3", map3)

}
