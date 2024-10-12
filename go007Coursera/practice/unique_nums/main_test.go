package main

import (
	"testing"
)

func TestReturnTheSame(t *testing.T) {
	testCases := []struct{
		input int
		expected int
	}{
		{1,1},
		{2,2},
		{3,3},
	}

	for _, tc := range(testCases) {
		result := ReturnTheSame(tc.input)
		if result != tc.expected {
            t.Errorf("ReturnTheSame(%d) = %d; expected %d", tc.input, result, tc.expected)
        }
	}
}