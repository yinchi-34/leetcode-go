package solution

import "testing"

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{100, 4, 200, 1, 3, 2},
			expected: 4,
		},
		{
			name:     "Example 2",
			nums:     []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			expected: 9,
		},
		{
			name:     "Example 3",
			nums:     []int{1, 0, 1, 2},
			expected: 3,
		},
		{
			name:     "Empty array",
			nums:     []int{},
			expected: 0,
		},
		{
			name:     "Single element",
			nums:     []int{7},
			expected: 1,
		},
		{
			name:     "All duplicates",
			nums:     []int{5, 5, 5},
			expected: 1,
		},
		{
			name:     "Negative numbers",
			nums:     []int{-1, -3, -2, 1, 0},
			expected: 5,
		},
		{
			name:     "Two runs, longer one second",
			nums:     []int{10, 11, 1, 2, 3, 4},
			expected: 4,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := longestConsecutive(test.nums)
			if result != test.expected {
				t.Errorf("longestConsecutive(%v) = %d, want %d", test.nums, result, test.expected)
			}
		})
	}
}
