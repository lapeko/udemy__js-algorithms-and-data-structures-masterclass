package main

func countZeroes(array []int) int {
	var (
		left   = 0
		right  = len(array) - 1
		middle = 0
	)

	for left <= right {
		middle = (left + right) / 2
		if array[middle] == 1 {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}

	if array[middle] == 1 {
		middle++
	}

	return len(array) - middle
}
