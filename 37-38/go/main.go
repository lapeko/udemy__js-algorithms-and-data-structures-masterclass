package main

import "fmt"

func quickSort[T comparable](a []T, comparator func(a, b T) int) []T {
	if comparator == nil {
		comparator = func(a, b T) int {
			return any(a).(int) - any(b).(int)
		}
	}

	var sort func(low, high int)

	sort = func(low, high int) {
		if high <= low {
			return
		}
		pivot := partition(a, comparator, low, high)
		fmt.Println(low, pivot-1, pivot+1, high)
		sort(low, pivot-1)
		sort(pivot+1, high)
	}

	sort(0, len(a)-1)

	return a
}

// TODO required
func partition[T comparable](a []T, comparator func(a, b T) int, low, high int) int {
	pivot := a[high]

	for j := low; j < high; j++ {
		if comparator(a[j], pivot) < 0 {
			a[j], a[low] = a[low], a[j]
			low++
		}
	}

	a[low], a[high] = a[high], a[low]

	return low
}
