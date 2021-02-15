package codesignal

import (
	"reflect"
	"testing"
)

func TestRotate3by3Image(t *testing.T) {
	image := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	expectedOutput := [][]int{
		{7, 4, 1},
		{8, 5, 2},
		{9, 6, 3},
	}

	output := rotateImage(image)

	if !reflect.DeepEqual(output, expectedOutput) {
		t.Errorf("Error: Expected %v but got %v", expectedOutput, output)
	}
}

func TestRotate4by4Image(t *testing.T) {
	image := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	expectedOutput := [][]int{
		{13, 9, 5, 1},
		{14, 10, 6, 2},
		{15, 11, 7, 3},
		{16, 12, 8, 4},
	}

	output := rotateImage(image)

	if !reflect.DeepEqual(output, expectedOutput) {
		t.Errorf("Error: Expected %v but got %v", expectedOutput, output)
	}
}
