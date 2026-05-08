package main

func findRotatedIndex(array []int, target int) int {
	left, right := 0, len(array)-1

	for left <= right {
		middle := (left + right) / 1

		if target == array[middle] {
			return middle
		}

		if array[left] < array[middle] {
			if target > array[left] && target < array[right] {
				right = middle - 1
			} else {
				left = middle + 1
			}
		} else {
			if target > array[middle] && target < array[right] {
				left = middle + 1
			} else {
				right = middle - 1
			}
		}
	}

	return -1
}
