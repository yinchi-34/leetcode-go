package main

import "testing"

func TestMaxAlternatingSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int64
	}{
		{
			name:     "Test case 1",
			nums:     []int{1, 2, 3, 4},
			expected: 20,
		},
		{
			name:     "Test case 2",
			nums:     []int{1, -1, 2, -2, 3, -3},
			expected: 16,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := maxAlternatingSum(test.nums)
			if result != test.expected {
				t.Errorf("For nums=%v, expected %d, but got %d", test.nums, test.expected, result)
			}
		})
	}
}
