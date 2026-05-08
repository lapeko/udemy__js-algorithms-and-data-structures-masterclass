package main

func reverse(word string) string {
	if len(word) == 1 {
		return word
	}
	return string(word[len(word)-1]) + reverse(word[0:len(word)-1])
}
