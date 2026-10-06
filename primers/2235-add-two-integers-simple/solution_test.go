package main

import "testing"

func TestSum(t *testing.T) {
	tests := []struct {
		name string
		num1 int
		num2 int
		expected int
	}{
		{
			name:"positive numbers",
			num1: 5,
			num2: 10,
			expected: 15,
		},
		{
			name: "negative numbers",
			num1: -5,
			num2: -10,
			expected: -15,
		},
		{
			name: "mixed numbers",
			num1: -5,
			num2: 10,
			expected: 5,
		},
		{
			name: "zero",
			num1: 0,
			num2: 0,
			expected: 0,
		},
	}



	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := sum(test.num1, test.num2)

			if result != test.expected {
				t.Errorf("sum(%d, %d) = %d; expected %d", test.num1, test.num2, result, test.expected)
			}
		})
	}
}


// To run the tests, use the command: go test -v

// table-driven tests are a common pattern in Go for testing multiple scenarios with different inputs and expected outputs. Each test case is defined in a struct, and the tests are executed in a loop, allowing for easy addition of new test cases.