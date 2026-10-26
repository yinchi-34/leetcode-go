package solution

func peakvalleyDifference(nums []int) int {
	n := len(nums)
	res := 0

	for i := 1; i < n-1; i++ {
		if nums[i] < nums[i-1] && nums[i] < nums[i+1] {
			res -= nums[i]
		}
		if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
			res += nums[i]
		}
	}
	return res
}
