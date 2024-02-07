package main

import "testing"

func TestSum(t *testing.T) {
	t.Run("Test 1", func(t *testing.T) {
		result := Sum(2, 2)
		if result != 4 {
			t.Error("Expected 4, got ", result)
		}
	})

	t.Run("Test 2", func(t *testing.T) {
		result := Sum(222, 2)
		if result != 224 {
			t.Error("Expected 224, got ", result)
		}
	})

	t.Run("Test 3", func(t *testing.T) {
		result := Sum(-2, 3)
		if result != 1 {
			t.Error("Expected 1, got ", result)
		}
	})
}

// go test -v