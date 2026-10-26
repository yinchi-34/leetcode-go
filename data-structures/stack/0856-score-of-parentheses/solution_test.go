package solution

import (
	"strings"
	"testing"
)

func TestScoreOfParentheses(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected int
	}{
		{
			name:     "single_pair",
			s:        "()",
			expected: 1,
		},
		{
			name:     "nesting_doubles",
			s:        "(())",
			expected: 2,
		},
		{
			name:     "concatenation_adds",
			s:        "()()",
			expected: 2,
		},
		{
			name:     "nesting_and_concatenation",
			s:        "(()(()))",
			expected: 6,
		},
		{
			name:     "max_depth",
			s:        strings.Repeat("(", 25) + strings.Repeat(")", 25),
			expected: 1 << 24,
		},
		{
			name:     "max_length_flat",
			s:        strings.Repeat("()", 25),
			expected: 25,
		},
	}

	solutions := []struct {
		name string
		fn   func(string) int
	}{
		{"stack", scoreOfParentheses},
		{"depth", scoreOfParenthesesDepth},
	}

	for _, sol := range solutions {
		for _, test := range tests {
			t.Run(sol.name+"/"+test.name, func(t *testing.T) {
				res := sol.fn(test.s)
				if res != test.expected {
					t.Errorf("%s(%q) = %d, want %d", sol.name, test.s, res, test.expected)
				}
			})
		}
	}
}
