package solution

import "testing"

func TestSortPermutation(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		excepted int
	}{
		{
			name:     "identity_permutation",
			nums:     []int{0, 1, 2, 3},
			excepted: 0,
		},
		{
			name:     "nonempty_support",
			nums:     []int{0, 1, 2, 3, 4, 6, 7, 5},
			excepted: 4,
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
