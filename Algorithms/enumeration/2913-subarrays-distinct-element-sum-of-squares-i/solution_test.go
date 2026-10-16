package solution

import "testing"

func TestSumCounts(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "test1",
			nums:     []int{1, 2, 1},
			expected: 15,
		},
		{
			name:     "test2",
			nums:     []int{1, 1},
			expected: 3,
		},
		{
			name:     "single element",
			nums:     []int{5},
			expected: 1,
		},
		{
			name:     "all distinct",
			nums:     []int{1, 2, 3},
			expected: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumCounts(tt.nums); got != tt.expected {
				t.Errorf("sumCounts() = %v, want %v", got, tt.expected)
			}
		})
	}
}
