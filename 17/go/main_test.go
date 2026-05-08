package main

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	tests := []struct {
		num      int
		expected int
	}{
		{
			num:      1,
			expected: 1,
		},
		{
			num:      2,
			expected: 2,
		},
		{
			num:      4,
			expected: 24,
		},
		{
			num:      7,
			expected: 5040,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			if res := factorial(tt.num); res != tt.expected {
				t.Errorf("factorial(%d) returns %d. Expected %d", tt.num, res, tt.expected)
			}
		})
	}
}
