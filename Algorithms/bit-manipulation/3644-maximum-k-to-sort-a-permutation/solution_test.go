package solution

import "testing"

func TestSortPermutation(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		excepted int
	}{
		{
			name: "single_element",
			nums: []int{
				0,
			},
			excepted: 0,
		},
		{
			name:     "already_sorted",
			nums:     []int{0, 1, 2, 3},
			excepted: 0,
		},
		{
			name:     "misplaced_zero",
			nums:     []int{4, 3, 2, 1, 0},
			excepted: 0,
		},
		{
			name:     "common_low_bit",
			nums:     []int{0, 3, 2, 1},
			excepted: 1,
		},
		{
			name:     "common_high_bit",
			nums:     []int{0, 1, 3, 2},
			excepted: 2,
		},
		{
			name:     "no_common_bits",
			nums:     []int{0, 2, 3, 1},
			excepted: 0,
		},
		{
			name:     "multiple_common_bits",
			nums:     []int{0, 1, 2, 7, 4, 5, 6, 3},
			excepted: 3,
		},
		{
			name:     "multiple_no_common_bits",
			nums:     []int{0, 1, 2, 3, 7, 5, 6, 4, 3},
			excepted: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := sortPermutation(test.nums)
			if result != test.excepted {
				t.Errorf("sortPermutation(%d) = %d, want %d", test.nums, result, test.excepted)
			}
		})
	}
}
