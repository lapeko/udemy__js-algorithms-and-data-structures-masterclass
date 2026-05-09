package main

func someRecursive(array []int, callback func(int) bool) bool {
	if len(array) == 0 {
		return false
	}
	return callback(array[0]) || someRecursive(array[1:], callback)
}
