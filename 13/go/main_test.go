package main

import "testing"

func TestMain(t *testing.T) {
	var tests = []struct {
		name   string
		array  []int
		expect int
	}{
		{
			name:   "1",
			array:  []int{1, 1, 1, 1, 0, 0},
			expect: 2,
		},
		{
			name:   "2",
			array:  []int{1, 0, 0, 0, 0},
			expect: 4,
		},
		{
			name:   "3",
			array:  []int{0, 0, 0},
			expect: 3,
		},
		{
			name:   "4",
			array:  []int{1, 1, 1, 1},
			expect: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res := countZeroes(tt.array); res != tt.expect {
				t.Errorf("countZeroes(%v). got %d when %d expected", tt.array, res, tt.expect)
			}
		})
	}
}
