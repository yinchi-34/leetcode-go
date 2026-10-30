package solution

import (
	"slices"
	"testing"
)

// buildTree builds a tree from LeetCode's level-order form, where nil marks a missing child.
func buildTree(vals []any) *TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}
	root := &TreeNode{Val: vals[0].(int)}
	queue := []*TreeNode{root}
	i := 1
	for len(queue) > 0 && i < len(vals) {
		node := queue[0]
		queue = queue[1:]
		if i < len(vals) && vals[i] != nil {
			node.Left = &TreeNode{Val: vals[i].(int)}
			queue = append(queue, node.Left)
		}
		i++
		if i < len(vals) && vals[i] != nil {
			node.Right = &TreeNode{Val: vals[i].(int)}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

func TestInorderTraversal(t *testing.T) {
	tests := []struct {
		name     string
		root     []any
		expected []int
	}{
		{
			name:     "Example 1",
			root:     []any{1, nil, 2, 3},
			expected: []int{1, 3, 2},
		},
		{
			name:     "Example 2",
			root:     []any{1, 2, 3, 4, 5, nil, 8, nil, nil, 6, 7, 9},
			expected: []int{4, 2, 6, 5, 7, 1, 3, 9, 8},
		},
		{
			name:     "Empty tree",
			root:     []any{},
			expected: []int{},
		},
		{
			name:     "Single node",
			root:     []any{1},
			expected: []int{1},
		},
		{
			name:     "Left-skewed chain",
			root:     []any{3, 2, nil, 1},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Binary search tree gives sorted order",
			root:     []any{4, 2, 6, 1, 3, 5, 7},
			expected: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name:     "Negative values",
			root:     []any{0, -1, 1},
			expected: []int{-1, 0, 1},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := inorderTraversal(buildTree(test.root))
			if !slices.Equal(result, test.expected) {
				t.Errorf("inorderTraversal(%v) = %v, want %v", test.root, result, test.expected)
			}
		})
	}
}
