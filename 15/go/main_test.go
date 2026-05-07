package main

import "testing"

func TestFindRotatedIndex(t *testing.T) {
	tests := []struct {
		name   string
		array  []int
		target int
		expect int
	}{
		{
			name:   "Test 1",
			array:  []int{3, 4, 1, 2},
			target: 4,
			expect: 1,
		},
		{
			name:   "Test 2",
			array:  []int{6, 7, 8, 9, 1, 2, 3, 4},
			target: 8,
			expect: 2,
		},
		{
			name:   "Test 3",
			array:  []int{6, 7, 8, 9, 1, 2, 3, 4},
			target: 3,
			expect: 6,
		},
		{
			name:   "Test 4",
			array:  []int{37, 44, 66, 102, 10, 22},
			target: 14,
			expect: -1,
		},
		{
			name:   "Test 5",
			array:  []int{6, 7, 8, 9, 1, 2, 3, 4},
			target: 12,
			expect: -1,
		}, {
			name:   "Test 6",
			array:  []int{11, 12, 13, 14, 15, 16, 3, 5, 7, 9},
			target: 16,
			expect: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res := findRotatedIndex(tt.array, tt.target); res != tt.expect {
				t.Errorf("findRotatedIndex(%v, %d) returns %d. Expected %d", tt.array, tt.target, res, tt.expect)
			}
		})
	}
}
