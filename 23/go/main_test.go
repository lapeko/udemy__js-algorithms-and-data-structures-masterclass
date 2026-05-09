package main

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	isOdd := func(num int) bool {
		return num%2 == 1
	}

	tests := []struct {
		array    []int
		name     string
		callback func(int) bool
		expected bool
	}{
		{[]int{1, 2, 3, 4}, "is odd", isOdd, true},
		{[]int{4, 6, 8, 9}, "is odd", isOdd, true},
		{[]int{4, 6, 8}, "is odd", isOdd, false},
		{
			array: []int{4, 6, 8},
			name:  "above ten",
			callback: func(num int) bool {
				return num > 10
			},
			expected: false,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test %d. Callback %s", i+1, tt.name), func(t *testing.T) {
			if res := someRecursive(tt.array, tt.callback); res != tt.expected {
				t.Errorf("someRecursive(%v, callback) return %t when %t expected", tt.array, res, tt.expected)
			}
		})
	}
}
