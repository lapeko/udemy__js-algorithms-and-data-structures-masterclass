package main

func flatten(array []any) []int {
	res := []int{}

	for _, item := range array {
		switch v := item.(type) {
		case int:
			res = append(res, v)
		case []any:
			res = append(res, flatten(v)...)
		}
	}

	return res
}
