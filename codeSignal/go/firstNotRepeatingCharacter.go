package codesignal

func firstNotRepeatingCharacter(s string) string {
	duplicates := make(map[rune]int)

	for _, letter := range s {
		duplicates[letter] = duplicates[letter] + 1
	}

	for _, letter := range s {
		if duplicates[letter] == 1 {
			return string(letter)
		}
	}
	return "_"
}
