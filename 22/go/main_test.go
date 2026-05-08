package main

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	tests := []struct {
		word   string
		expect bool
	}{
		{"awesome", false},
		{"foobar", false},
		{"tacocat", true},
		{"amanaplanacanalpanama", true},
		{"amanaplanacanalpandemonium", false},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			if res := isPalindrome(tt.word); res != tt.expect {
				t.Errorf("reverse(%s) returns \"%t\" when \"%t\" expected", tt.word, res, tt.expect)
			}
		})
	}
}
