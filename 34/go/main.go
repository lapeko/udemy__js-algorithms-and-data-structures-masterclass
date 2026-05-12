package main

func insertionSort[T comparable](array []T, comparator func(a, b T) int) []T {
	if comparator == nil {
		comparator = func(a, b T) int {
			aa := any(a).(int)
			bb := any(b).(int)
			return aa - bb
		}
	}
	for i := 1; i < len(array); i++ {
		insertVal := array[i]
		insertIdx := i
		for j := i - 1; j >= 0; j-- {
			if comparator(array[j], insertVal) > 0 {
				array[j+1] = array[j]
				insertIdx = j
			} else {
				break
			}
		}
		array[insertIdx] = insertVal
	}
	return array
}
