package solution

import "testing"


func TestSmallestEvenMultiple(t *testing.T) {
	tests := []struct {
		name string
		n    int
		expected int
	}{
		{
			name:     "Test case 1",
			n:        5,
			expected: 10,
		},
		{
			name:     "Test case 2",
			n:        6,
			expected: 6,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := smallestEvenMultiple(test.n)
			if result != test.expected {
				t.Errorf("For n=%d, expected %d, but got %d", test.n, test.expected, result)
			}
		})
	}
}