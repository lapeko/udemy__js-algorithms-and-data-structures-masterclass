package main

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	tests := []struct {
		obj      map[string]any
		expected int
	}{
		{
			obj: map[string]any{
				"outer": 2,
				"obj": map[string]any{
					"inner": 2,
					"otherObj": map[string]any{
						"superInner":     2,
						"notANumber":     true,
						"alsoNotANumber": "yup",
					},
				},
			},
			expected: 6,
		},
		{
			obj: map[string]any{
				"a": 2,
				"b": map[string]any{
					"b": 2,
					"bb": map[string]any{
						"b": 3,
						"bb": map[string]any{
							"b": 2,
						},
					},
				},
				"c": map[string]any{
					"c": map[string]any{
						"c": 2,
					},
					"cc":  "ball",
					"ccc": 5,
				},
				"d": 1,
				"e": map[string]any{
					"e": map[string]any{
						"e": 2,
					},
					"ee": "car",
				},
			},
			expected: 10,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			if res := nestedEvenSum(tt.obj); res != tt.expected {
				t.Errorf("nestedEvenSum returns %d when %d expected", res, tt.expected)
			}
		})
	}
}
