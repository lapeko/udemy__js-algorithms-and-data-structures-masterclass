package main

import "testing"

func TestMain(t *testing.T) {
	tests := []struct {
		name     string
		array    []int
		search   int
		expected int
	}{
		{name: "1", array: []int{1, 1, 2, 2, 2, 2, 3}, search: 2, expected: 4},
		{name: "1", array: []int{1, 1, 2, 2, 2, 2, 3}, search: 3, expected: 1},
		{name: "1", array: []int{1, 1, 2, 2, 2, 2, 3}, search: 1, expected: 2},
		{name: "1", array: []int{1, 1, 2, 2, 2, 2, 3}, search: 4, expected: -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res := sortedFrequency(tt.array, tt.search); res != tt.expected {
				t.Errorf("sortedFrequency(%v, %d). Got: %d. Expected: %d", tt.array, tt.search, res, tt.expected)
			}
		})
	}
}
