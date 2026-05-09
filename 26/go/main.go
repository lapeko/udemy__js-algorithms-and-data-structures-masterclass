package main

func nestedEvenSum(obj map[string]any) int {
	sum := 0
	for _, val := range obj {
		switch v := val.(type) {
		case int:
			if v%2 == 0 {
				sum += v
			}
		case map[string]any:
			sum += nestedEvenSum(v)
		}
	}
	return sum
}
