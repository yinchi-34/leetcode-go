package main

import "sort"

func smallestEvenMultiple(nums []int) int64 {
	n := len(nums)

	for i := 0; i < n; i++ {
		nums[i] *= nums[i]
	}
	sort.Ints(nums)

	var res int64
	for i := 0; i < n; i++ {
		if i >= n/2 {
			res += int64(nums[i])
		} else {
			res -= int64(nums[i])
		}
	}
	return int64(res)
}
