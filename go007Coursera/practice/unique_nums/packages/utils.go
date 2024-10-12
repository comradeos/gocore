package utils

import "fmt"

func HelloUtils() {
	fmt.Println("HelloUtils")
}

func ContainsInt(values []int, target int) bool {
	for _, v := range(values) {
		if (v == target) {
			return true
		}
 	}

	return false
}

func GetUniqueNumbers(values []int) []int {
	result := []int {}

	for _, v := range(values) {
		contains := ContainsInt(result, v)

		if (contains) {
			continue
		}

		result = append(result, v)
	}

	return result
}