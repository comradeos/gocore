package main

import "testing"

func TestMultple(t *testing.T) {

	t.Run("Test 1", func(t *testing.T) {
		result := Multple(2, 2)
		if result != 6 {
			t.Error("Expected 6, got ", result)
		}
	}

	t.Run("Test 2", func(t *testing.T) {
		result := Multple(234, 234)
		if result != 54756 {
			t.Error("Expected 54756, got ", result)
		}
	}

	t.Run("Test 3", func(t *testing.T) {
		result := Multple(-2, 3)
		if result != -6 {
			t.Error("Expected -6, got ", result)
		}
	}
}