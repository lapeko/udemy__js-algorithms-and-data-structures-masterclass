package main

import "testing"

func TestMain(t *testing.T) {
	tests := []struct {
		name     string
		base     int
		pow      int
		expected int
	}{
		{
			name:     "Test 1",
			base:     2,
			pow:      0,
			expected: 1,
		},
		{
			name:     "Test 2",
			base:     2,
			pow:      2,
			expected: 4,
		},
		{
			name:     "Test 3",
			base:     2,
			pow:      4,
			expected: 16,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res := power(tt.base, tt.pow); res != tt.expected {
				t.Errorf("power(%d, %d) returns %d when expected %d", tt.base, tt.pow, res, tt.expected)
			}
		})
	}
}
