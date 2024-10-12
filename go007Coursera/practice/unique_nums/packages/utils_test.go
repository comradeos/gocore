package utils

import (
	"reflect"
	"testing"
)

// Тест для функции GetUniqueNumbers
func TestGetUniqueNumbers(t *testing.T) {
	testCases := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 1, 2, 2, 3, 3}, []int{1, 2, 3}},
		{[]int{4, 4, 4, 4}, []int{4}},
		{[]int{10, 20, 10, 30, 20}, []int{10, 20, 30}},
		{[]int{}, []int{}}, // Пустой срез
		{[]int{7, 7, 7, 8, 9}, []int{7, 8, 9}},
	}

	for _, tc := range testCases {
		result := GetUniqueNumbers(tc.input)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("GetUniqueNumbers(%v) = %v; expected %v", tc.input, result, tc.expected)
		}
	}
}
