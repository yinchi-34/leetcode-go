package main

import "testing"

func TestSumOfSquares(t *testing.T) {
	tests := []struct {
		name string	
		nums []int
		expected int
	}{
		{
			name: "test1",
			nums: []int{1, 2, 3, 4},
			expected: 21,
		},
		{
			name: "test2",
			nums: []int{2, 7, 1, 19, 18, 3},
			expected: 63,
		},
		{
			name: "test3",
			nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: 91,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfSquares(tt.nums); got != tt.expected {
				t.Errorf("sumOfSquares() = %v, want %v", got, tt.expected)
			}
		})
	}
}