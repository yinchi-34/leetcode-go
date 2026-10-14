package main

import "testing"


func TestMinInsertions(t *testing.T) {
	tests := []struct {
		name	 string
		s 	  string
		expected int
	} {
		{
			name:     "Test 1",
			s:        "())",
			expected: 0,
		},
		{
			name:     "Test 2",
			s:        "(((",
			expected: 6,
		},
		{
			name:     "Test 3",
			s:        "()",
			expected: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := minAddToMakeValid(test.s)
			if result != test.expected {
				t.Errorf("minAddToMakeValid(%s) = %d, want %d", test.s, result, test.expected)
			}
	})
}
}
