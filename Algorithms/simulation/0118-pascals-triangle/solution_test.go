package solution

import (
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	tests := []struct {
		name     string
		numRows  int
		expected [][]int
	}{
		{
			name:    "1 row",
			numRows: 1,
			expected: [][]int{
				{1},
			},
		},
		{
			name:    "5 rows",
			numRows: 5,
			expected: [][]int{
				{1},
				{1, 1},
				{1, 2, 1},
				{1, 3, 3, 1},
				{1, 4, 6, 4, 1},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := generate(test.numRows)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("generate(%d) = %v, want %v", test.numRows, result, test.expected)
			}
		})
	}
}
