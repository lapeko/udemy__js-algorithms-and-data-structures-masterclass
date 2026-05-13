package main

import (
	"slices"
	"testing"
)

const funcName = "quickSort"

func sortFunc[T comparable](array []T, comparator func(a, b T) int) []T {
	return quickSort(array, comparator)
}

func check[T comparable](t *testing.T, array []T, comparator func(a, b T) int, expected []T) {
	t.Helper()
	if res := sortFunc(slices.Clone(array), comparator); !slices.Equal(res, expected) {
		t.Errorf("%s(%v) returns %v when %v expected", funcName, array, res, expected)
	}
}

func TestMain(t *testing.T) {
	strComp := func(a, b string) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	}

	type testObj struct {
		name string
		age  int
	}

	oldestToYoungest := func(a, b testObj) int {
		return b.age - a.age
	}

	moarKittyData := []testObj{{
		name: "LilBub",
		age:  7,
	}, {
		name: "Garfield",
		age:  40,
	}, {
		name: "Heathcliff",
		age:  45,
	}, {
		name: "Blue",
		age:  1,
	}, {
		name: "Grumpy",
		age:  6,
	}}

	expectedMoarKittyData := []testObj{{
		name: "Heathcliff",
		age:  45,
	}, {
		name: "Garfield",
		age:  40,
	}, {
		name: "LilBub",
		age:  7,
	}, {
		name: "Grumpy",
		age:  6,
	}, {
		name: "Blue",
		age:  1,
	}}

	t.Run("Numbers", func(t *testing.T) {
		check(t, []int{4, 20, 12, 10, 7, 9}, nil, []int{4, 7, 9, 10, 12, 20})
		check(t, []int{0, -10, 7, 4}, nil, []int{-10, 0, 4, 7})
		check(t, []int{1, 2, 3}, nil, []int{1, 2, 3})
		check(t, []int{}, nil, []int{})
		check(t, []int{4, 3, 5, 3, 43, 232, 4, 34, 232, 32, 4, 35, 34, 23, 2, 453, 546, 75, 67, 4342, 32}, nil, []int{2, 3, 3, 4, 4, 4, 5, 23, 32, 32, 34, 34, 35, 43, 67, 75, 232, 232, 453, 546, 4342})
	})

	t.Run("Strings", func(t *testing.T) {
		check(t, []string{"LilBub", "Garfield", "Heathcliff", "Blue", "Grumpy"}, strComp, []string{"Blue", "Garfield", "Grumpy", "Heathcliff", "LilBub"})
	})

	t.Run("Objects", func(t *testing.T) {
		check(t, moarKittyData, oldestToYoungest, expectedMoarKittyData)
	})
}
