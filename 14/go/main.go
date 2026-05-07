package main

import (
	"math"
)

func sortedFrequency(array []int, search int) int {
	min := findMinIdx(array, search)
	if min == -1 {
		return -1
	}
	return findMaxIdx(array, search) - min + 1
}

func findMinIdx(array []int, search int) int {
	var left, right, idx = 0, len(array) - 1, math.MaxInt

	if array[left] == search {
		return 0
	}

	for left <= right {
		middle := (left + right) / 2
		if array[middle] == search {
			idx = middle
		}
		if array[middle] < search {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}

	if idx != math.MaxInt {
		return idx
	}

	return -1
}

func findMaxIdx(array []int, search int) int {
	var left, right, idx = 0, len(array) - 1, math.MaxInt

	if array[right] == search {
		return len(array) - 1
	}

	for left <= right {
		middle := (left + right) / 2
		if array[middle] == search {
			idx = middle
		}
		if array[middle] > search {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}

	if idx != math.MaxInt {
		return idx
	}

	return -1
}
