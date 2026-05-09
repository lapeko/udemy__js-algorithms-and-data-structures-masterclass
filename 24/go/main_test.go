package main

import (
	"fmt"
	"slices"
	"testing"
)

func TestMain(t *testing.T) {
	tests := []struct {
		array    []any
		expected []int
	}{
		{array: []any{1, 2, 3, []any{4, 5}}, expected: []int{1, 2, 3, 4, 5}},
		{array: []any{1, []any{2, []any{3, 4}, []any{[]any{5}}}}, expected: []int{1, 2, 3, 4, 5}},
		{array: []any{[]any{1}, []any{2}, []any{3}}, expected: []int{1, 2, 3}},
		{array: []any{[]any{[]any{[]any{1}, []any{[]any{2}}, []any{[]any{[]any{[]any{[]any{[]any{3}}}}}}}}}, expected: []int{1, 2, 3}},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			if res := flatten(tt.array); !slices.Equal(res, tt.expected) {
				t.Errorf("flatten(%v) returns %v, when expected %v", tt.array, res, tt.expected)
			}
		})
	}
}
