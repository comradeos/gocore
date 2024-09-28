package chapters

import (
	"fmt"
)

func Ch003ControlStructures() {
	fmt.Println("============== Hello from Ch003003ControlStructures() ==============")

	// if

	var num = 10
	if num > 5 {
		fmt.Println("num > 5")
	}

	var map1 = map[string]int{"key1": 1, "key2": 2}
	if val, exists := map1["key3"]; exists {
		fmt.Println("val", val)
	} else {
		fmt.Println("key3 does not exist")
	}

	// switch

	var num1 = 1
	switch num1 {
	case 1:
		fmt.Println("num1 == 1")
		fallthrough
	case 2:
		fmt.Println("num1 == 2")
	default:
		fmt.Println("num1 != 1 and num1 != 2")
	}

	// выход из внешнего цикла

OuterLoop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				break OuterLoop
			}
			fmt.Println("i", i, "j", j)
		}
	}

	// циклы for

	// for
	fmt.Println("for")
	for i := 0; i < 3; i++ {
		fmt.Println("i", i)
	}

	// for без условия
	fmt.Println("for without condition")
	var i = 0
	for {
		if i >= 3 {
			break
		}
		fmt.Println("i", i)
		i++
	}

	// for range
	fmt.Println("for range")
	var slice = []int{1, 2, 3}
	for index, value := range slice {
		fmt.Println("index", index, "value", value)
	}

	// for range map
	fmt.Println("for range map")
	var map2 = map[string]int{"key1": 1, "key2": 2}
	for key, value := range map2 {
		fmt.Println("key", key, "value", value)
	}

	// for range string
	fmt.Println("for range string")
	var str = "Hello, World!"
	for index, value := range str {
		fmt.Println("index", index, "value", string(value))
	}

}
