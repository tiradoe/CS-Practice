package codesignal

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	inputString := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	expectedOutput := len([]int{0, 1, 2, 3, 4})

	output := removeDuplicates(&inputString)

	if expectedOutput != output {
		t.Errorf("Error: Expected %d but got %d", expectedOutput, output)
	}
}
