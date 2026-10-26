package solution

import "testing"

func TestScoreOfParentheses(t *testing.T) {
	tests := []struct {
		name     	string
		s			string
		expected 	int
	}{
		{
			name: "Example 1", 
			s: "()",
			expected: 1,},
		{
			name: "Example 2", 
			s: "(())",
			expected: 2,
		},
		{
			name: "Example 3", 
			s: "()()",
			expected: 2,
		},
		{
			name: "Example 4", 
			s: "(()(()))",
			expected: 6,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := scoreofParentheses(test.s)
			if result != test.expected {
				t.Errorf("scoreofParentheses(%s) = %d, want %d", test.s, result, test.expected)
			}
		})
	}
}