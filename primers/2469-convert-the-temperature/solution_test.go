package main	

import "testing"

func TestConvertTemperature(t *testing.T) {
	tests := []struct {
		name     string
		celsius  float64
		expected []float64
	}{
		{
			name:     "Example 1",
			celsius:  36.50,
			expected: []float64{309.65, 97.70},
		},
		{
			name:     "Example 2",
			celsius:  122.11,
			expected: []float64{395.26, 251.80},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := convertTemperature(test.celsius)

			if len(result) != len(test.expected) {
				t.Errorf("convertTemperature(%f) = %v; expected %v", test.celsius, result, test.expected)
				return
			}
	})
}
}

