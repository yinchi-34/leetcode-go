package solution

import (
	"testing"

	iterative "leetcode-go/Algorithms/graph-search/3619-count-islands-with-total-value-divisible-by-k/iterative-dfs"
	recursive "leetcode-go/Algorithms/graph-search/3619-count-islands-with-total-value-divisible-by-k/recursive-dfs"
)

// cloneGrid 返回 grid 的深拷贝。递归解法会把访问过的格子置 0，
// 每次调用前复制一份，避免用例之间互相影响。
func cloneGrid(grid [][]int) [][]int {
	g := make([][]int, len(grid))
	for i, row := range grid {
		g[i] = append([]int(nil), row...)
	}
	return g
}

func fullGrid(m, n, v int) [][]int {
	g := make([][]int, m)
	for i := range g {
		g[i] = make([]int, n)
		for j := range g[i] {
			g[i][j] = v
		}
	}
	return g
}

func TestCountIslands(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		k        int
		expected int
	}{
		{
			name: "example_1",
			grid: [][]int{
				{0, 2, 1, 0, 0},
				{0, 5, 0, 0, 5},
				{0, 0, 1, 0, 0},
				{0, 1, 4, 7, 0},
				{0, 2, 0, 0, 8},
			},
			k:        5,
			expected: 2,
		},
		{
			// 对角线不算相连：若误把对角线当相邻，6 个 3 会连成一座总和 18 的岛，得到 1
			name: "example_2_diagonal_not_connected",
			grid: [][]int{
				{3, 0, 3, 0},
				{0, 3, 0, 3},
				{3, 0, 3, 0},
			},
			k:        3,
			expected: 6,
		},
		{
			name:     "all_water",
			grid:     [][]int{{0, 0}, {0, 0}},
			k:        1,
			expected: 0,
		},
		{
			name:     "single_cell_not_divisible",
			grid:     [][]int{{7}},
			k:        2,
			expected: 0,
		},
		{
			// 岛屿内有环：同一个格子可以从两个邻居到达。
			// 若显式栈在出栈时才标记且不检查，(1,0) 会被压栈两次，总和算成 5
			name:     "square_island_has_cycle",
			grid:     [][]int{{1, 1}, {1, 1}},
			k:        4,
			expected: 1,
		},
		{
			// 一条从 (0,0) 出发的螺旋形路径，依次需要向右、下、左、上、右走，
			// 少写任何一个方向都会把岛屿拆成几块，每块总和都小于 14
			name: "spiral_island_needs_all_directions",
			grid: [][]int{
				{1, 1, 1, 1},
				{0, 0, 0, 1},
				{1, 1, 0, 1},
				{1, 0, 0, 1},
				{1, 1, 1, 1},
			},
			k:        14,
			expected: 1,
		},
		{
			// m*n 与格子值都取上限：总和 1e11 超出 int32（截断后不能被 1e6 整除，会得到 0），
			// 递归解法的递归深度也会达到 1e5
			name:     "full_grid_max_sum",
			grid:     fullGrid(100, 1000, 1_000_000),
			k:        1_000_000,
			expected: 1,
		},
	}

	solutions := []struct {
		name string
		fn   func([][]int, int) int
	}{
		{"recursive_dfs", recursive.CountIslands},
		{"iterative_dfs", iterative.CountIslands},
	}

	for _, s := range solutions {
		for _, test := range tests {
			t.Run(s.name+"/"+test.name, func(t *testing.T) {
				result := s.fn(cloneGrid(test.grid), test.k)
				if result != test.expected {
					t.Errorf("%s(%s, %d) = %d, want %d", s.name, test.name, test.k, result, test.expected)
				}
			})
		}
	}
}
