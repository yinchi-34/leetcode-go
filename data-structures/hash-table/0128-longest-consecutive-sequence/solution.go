package solution

func longestConsecutive(nums []int) int {
	res := 0
	numSet := map[int]bool{}
	for _, num := range nums {
		numSet[num] = true
	}

	for num := range numSet {
		if !numSet[num-1] {
			currentNum := num
			currentStreak := 1
			for numSet[currentNum+1] {
				currentNum++
				currentStreak++
			}
			if res <= currentStreak {
				res = currentStreak
			}
		}
	}
	return res
}
