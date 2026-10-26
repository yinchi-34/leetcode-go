package solution

import "testing"

func TestPeakValleyDifference(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		excepted int
	}{
		{
			name:     "postive numbers",
			nums:     []int{1, 5, 2, 4, 1, 4, 0},
			excepted: 10,
		},
		{
			name:     "negative numbers",
			nums:     []int{-1, -5, -6, -1, 0},
			excepted: 6,
		},
		{
			name:     "more valley",
			nums:     []int{6, 5, 5, 3, 2, 1, 0},
			excepted: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := peakvalleyDifference(tt.nums)
			if result != tt.excepted {
				t.Errorf("peakvalleyDifference(%d) = %v, want %d", tt.nums, result, tt.excepted)
			}
		})
	}
}
