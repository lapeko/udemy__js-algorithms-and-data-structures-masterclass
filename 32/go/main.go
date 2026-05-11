package main

func bubbleSort[T comparable](array []T, comparator func(a T, b T) int) []T {
	if comparator == nil {
		comparator = func(a, b T) int {
			first := any(a).(int)
			second := any(b).(int)
			return first - second
		}
	}
	for i := 1; i < len(array)-1; i++ {
		swap := false
		for j := 0; j < len(array)-i; j++ {
			if comparator(array[j], array[j+1]) > 0 {
				swap = true
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
		if !swap {
			break
		}
	}
	return array
}
