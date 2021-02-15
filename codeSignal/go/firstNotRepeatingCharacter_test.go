package main

import (
	"testing"
)

func TestFirstNotRepeatingCharacter(t *testing.T) {
	inputString := "abacabad"
	expectedOutput := "c"

	output := firstNotRepeatingCharacter(inputString)

	if expectedOutput != output {
		t.Errorf("Error: Expected %s but got %s", expectedOutput, output)
	}
}

func TestFirstNotRepeatingCharacterNotFound(t *testing.T) {
	inputString := "aaaaaaaaaaaaaaaa"
	expectedOutput := "_"

	output := firstNotRepeatingCharacter(inputString)

	if expectedOutput != output {
		t.Errorf("Error: Expected %s but got %s", expectedOutput, output)
	}
}
