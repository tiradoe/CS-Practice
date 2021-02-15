package main

import (
	"testing"
)

func TestPalindromeCheck(t *testing.T) {
	inputString := "aabbaacc"
	expectedOutput := false

	output := checkPalindrome(inputString)

	if expectedOutput != output {
		t.Errorf("Error: Expected %t but got %t", expectedOutput, output)
	}
}
