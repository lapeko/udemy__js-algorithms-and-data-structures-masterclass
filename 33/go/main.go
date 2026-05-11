package main

func selectionSort[T comparable](array []T, comparator func(a, b T) int) []T {
	if comparator == nil {
		comparator = func(a, b T) int {
			aa := any(a).(int)
			bb := any(b).(int)
			return aa - bb
		}
	}
	for i := 0; i < len(array)-1; i++ {
		minIdx := i
		for j := i + 1; j < len(array); j++ {
			if comparator(array[j], array[minIdx]) < 0 {
				minIdx = j
			}
		}
		if minIdx != i {
			array[i], array[minIdx] = array[minIdx], array[i]
		}
	}
	return array
}
