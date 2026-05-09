package main

import (
	"fmt"
	"slices"
	"testing"
)

func TestMain(t *testing.T) {
	tests := []struct {
		words    []string
		expected []string
	}{
		{[]string{"car", "taco", "banana"}, []string{"Car", "Taco", "Banana"}},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			if res := capitalizeFirst(tt.words); !slices.Equal(res, tt.expected) {
				t.Errorf("capitalizeFirst(%v) returns %v when %v expected", tt.words, res, tt.expected)
			}
		})
	}
}
