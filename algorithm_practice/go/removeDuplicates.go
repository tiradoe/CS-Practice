package main

import "fmt"

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println("Before: ", nums)
	removeDuplicates(&nums)
}

func removeDuplicates(nums *[]int) int {
	var deduped []int
	for _, number := range *nums {
		if !find(number, deduped) {
			deduped = append(deduped, number)
		}
	}

	fmt.Println("After: ", deduped)
	nums = &deduped
	return len(*nums)
}

func find(number int, list []int) bool {
	for num := range list {
		if num == number {
			return true
		}
	}

	return false
}
