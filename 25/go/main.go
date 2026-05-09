package main

import "unicode"

func capitalizeFirst(words []string) []string {
	if len(words) == 0 {
		return []string{}
	}
	return append([]string{capitalize(words[0])}, capitalizeFirst(words[1:])...)
}

func capitalize(word string) string {
	if len(word) == 0 {
		return word
	}
	runes := []rune(word)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
