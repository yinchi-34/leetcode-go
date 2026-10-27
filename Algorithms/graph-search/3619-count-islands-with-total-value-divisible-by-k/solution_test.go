package solution

import "testing"

func cloneGrid(grid [][]int) [][]int {
	g := make([][]int, len(grid))
	for i, row := range grid {
		g[i] = append([]int(nil), row...)
	}
	return g
}

func TestCountIslands(t *testing.T) {
	longRow := make([]int, 100000)
	for i := range longRow {
		longRow[i] = 1_000_000
	}

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
			// 蛇形岛屿需要多次转向才能走完，只有搜完全部 14 格总和才是 14
			name: "snake_island",
			grid: [][]int{
				{1, 1, 1, 1},
				{0, 0, 0, 1},
				{1, 1, 1, 1},
				{1, 0, 0, 0},
				{1, 1, 1, 1},
			},
			k:        14,
			expected: 1,
		},
		{
			// m*n 取到上限 1e5、每格取到上限 1e6，总和 1e11 超出 int32；递归深度也达到 1e5
			name:     "long_row_max_sum",
			grid:     [][]int{longRow},
			k:        1_000_000,
			expected: 1,
		},
	}

	solutions := []struct {
		name string
		fn   func([][]int, int) int
	}{
		{"recursive_dfs", countIslands},
		{"stack_dfs", countIslandsStack},
	}

	for _, s := range solutions {
		for _, test := range tests {
			t.Run(s.name+"/"+test.name, func(t *testing.T) {
				result := s.fn(cloneGrid(test.grid), test.k)
				if result != test.expected {
					t.Errorf("%s(%v, %d) = %d, want %d", s.name, test.grid, test.k, result, test.expected)
				}
			})
		}
	}
}
