package main

func factorial(num int) int {
	if num < 1 {
		return 1
	}

	return num * factorial(num-1)
}

// func factorial(num int) int {
// 	sum := 1

// 	for i := num; i > 1; i-- {
// 		sum *= i
// 	}

// 	return sum
// }
