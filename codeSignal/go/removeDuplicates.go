package codesignal

func removeDuplicates(nums *[]int) int {
	var deduped []int
	for _, number := range *nums {
		if !find(number, deduped) {
			deduped = append(deduped, number)
		}
	}

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
