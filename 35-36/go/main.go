package main

import "fmt"

func mergeSort[T comparable](array []T, comparator func(a, b T) int) []T {
	if comparator == nil {
		comparator = func(a, b T) int {
			aa := any(a).(int)
			bb := any(b).(int)
			return aa - bb
		}
	}

	var fn func(a []T) []T

	fn = func(a []T) []T {
		if len(a) <= 1 {
			return a
		}

		middle := len(a) / 2

		return merge(
			fn(a[:middle]),
			fn(a[middle:]),
			comparator,
		)
	}

	return fn(array)
}

func merge[T comparable](a, b []T, comparator func(a, b T) int) []T {
	l, r, lr, size := 0, 0, 0, len(a)+len(b)
	c := make([]T, size)

	for lr < size {
		switch {
		case l >= len(a):
			c[lr] = b[r]
			r++
		case r >= len(b):
			c[lr] = a[l]
			l++
		case comparator(a[l], b[r]) < 0:
			c[lr] = a[l]
			l++
		default:
			c[lr] = b[r]
			r++
		}
		lr++
	}

	return c
}

func main() {
	fmt.Printf("%v", merge([]int{3, 4, 5}, []int{3, 6, 7}, func(a, b int) int {
		return a - b
	}))
}
