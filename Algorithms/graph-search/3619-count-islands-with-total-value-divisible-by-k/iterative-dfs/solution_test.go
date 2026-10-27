package solution

import (
	"testing"

	"leetcode-go/Algorithms/graph-search/3619-count-islands-with-total-value-divisible-by-k/testcases"
)

func TestCountIslands(t *testing.T) {
	for _, test := range testcases.Cases {
		t.Run(test.Name, func(t *testing.T) {
			result := countIslands(testcases.CloneGrid(test.Grid), test.K)
			if result != test.Expected {
				t.Errorf("countIslands(%s, %d) = %d, want %d", test.Name, test.K, result, test.Expected)
			}
		})
	}
}
