package main

import "testing"

func TestRemoveOuterParentheses(t *testing.T) {
	test := []struct {
		name	string
		s		string
		excepted string	
} {
	{
		name:   "Test 1",
		s:      "(()())(())",
		excepted: "()()()",
	},
	{
		name:   "Test 2",
		s:      "(()())(())(()(()))",
		excepted: "()()()()(())",
	},
	{
		name:   "Test 3",
		s:      "()()",
		excepted: "",
	},
}
	
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			res := removeOuterParentheses(tt.s)
			if res != tt.excepted {
				t.Errorf("removeOuterParentheses(%s) = %s; want %s", tt.s, res, tt.excepted)
			}
		})
	}
}