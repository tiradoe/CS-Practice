package main

//func main() {
//	fmt.Println("bananahammock: ", checkPalindrome("bananahammock"))
//	fmt.Println("racecar: ", checkPalindrome("racecar"))
//	fmt.Println("abac: ", checkPalindrome("abac"))
//}

func checkPalindrome(inputString string) bool {
	i := len(inputString) - 1

	for _, letter := range inputString {
		if string(letter) != string(inputString[i]) {
			return false
		}

		i--
	}

	return true
}
