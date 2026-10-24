package solution

func sortPermutation(nums []int) int {
	k := -1
	n := len(nums)

	for i := 0; i < n; i++ {
		if nums[i] == i {
			continue
		}
		k &= nums[i]
	}
	if k == -1 {
		return 0
	}
	return k
}
