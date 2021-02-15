package codesignal

func rotateImage(a [][]int) [][]int {
	// Transpose the grid
	// 1,2,3      1,4,7
	// 4,5,6  --> 2,5,8
	// 7,8,9      3,6,9

	for rowIndex := 0; rowIndex < len(a); rowIndex++ {
		for colIndex := rowIndex; colIndex < len(a); colIndex++ {
			a[colIndex][rowIndex], a[rowIndex][colIndex] = a[rowIndex][colIndex], a[colIndex][rowIndex]
		}
	}

	// Flip the sides:
	// 1,4,7    7,4,1
	// 2,5,8 -> 8,5,2
	// 3,6,9    9,6,3

	for colIndex := 0; colIndex < len(a)/2; colIndex++ {
		for rowIndex := 0; rowIndex < len(a); rowIndex++ {
			a[rowIndex][colIndex], a[rowIndex][len(a)-colIndex-1] = a[rowIndex][len(a)-colIndex-1], a[rowIndex][colIndex]
		}
	}
	return a
}
