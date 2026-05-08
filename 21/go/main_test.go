package main

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	tests := []struct {
		word   string
		expect string
	}{
		{word: "awesome", expect: "emosewa"},
		{word: "rithmschool", expect: "loohcsmhtir"},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			if res := reverse(tt.word); res != tt.expect {
				t.Errorf("reverse(%s) returns \"%s\" when \"%s\" expected", tt.word, res, tt.expect)
			}
		})
	}
}
