package main

import "testing"

func TestSum(t *testing.T) {
	res := Sum(10, 20)
	if res != 30 {
		t.Errorf("Expected 30, but got %d", res)
	}
}