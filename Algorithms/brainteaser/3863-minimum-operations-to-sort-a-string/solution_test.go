package solution

import "testing"

func TestMinOperations(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		excepted int
	}{
		{
			name:     "already_sorted",
			s:        "abcde",
			excepted: 0,
		},
		{
			name:     "one_operation",
			s:        "acbed",
			excepted: 1,
		},
		{
			name:     "two_operations",
			s:        "beadc",
			excepted: 2,
		},
		{
			name:     "three_operations",
			s:        "edcba",
			excepted: 3,
		},
		{
			name:     "impossible",
			s:        "ca",
			excepted: -1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := minOperations(test.s)
			if result != test.excepted {
				t.Errorf("minOperations(%s) = %d, want %d", test.s, result, test.excepted)
			}
		})
	}
}
