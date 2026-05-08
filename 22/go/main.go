package main

func isPalindrome(word string) bool {
	if len(word) < 2 {
		return true
	}
	return word[0] == word[len(word)-1] && isPalindrome(word[1:len(word)-1])
}
