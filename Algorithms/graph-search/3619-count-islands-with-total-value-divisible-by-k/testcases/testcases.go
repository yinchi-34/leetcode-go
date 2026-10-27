// Package testcases 保存 3619 两种解法共用的测试用例。
package testcases

// Case 是一组输入和预期输出。
type Case struct {
	Name     string
	Grid     [][]int
	K        int
	Expected int
}

// CloneGrid 返回 grid 的深拷贝。递归解法会把访问过的格子置 0，
// 每次调用前复制一份，避免用例之间互相影响。
func CloneGrid(grid [][]int) [][]int {
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

// Cases 是两种解法共用的测试用例。
var Cases = []Case{
	{
		Name: "example_1",
		Grid: [][]int{
			{0, 2, 1, 0, 0},
			{0, 5, 0, 0, 5},
			{0, 0, 1, 0, 0},
			{0, 1, 4, 7, 0},
			{0, 2, 0, 0, 8},
		},
		K:        5,
		Expected: 2,
	},
	{
		// 对角线不算相连：若误把对角线当相邻，6 个 3 会连成一座总和 18 的岛，得到 1
		Name: "example_2_diagonal_not_connected",
		Grid: [][]int{
			{3, 0, 3, 0},
			{0, 3, 0, 3},
			{3, 0, 3, 0},
		},
		K:        3,
		Expected: 6,
	},
	{
		Name:     "all_water",
		Grid:     [][]int{{0, 0}, {0, 0}},
		K:        1,
		Expected: 0,
	},
	{
		Name:     "single_cell_not_divisible",
		Grid:     [][]int{{7}},
		K:        2,
		Expected: 0,
	},
	{
		// 岛屿内有环：同一个格子可以从两个邻居到达。
		// 若显式栈在出栈时才标记且不检查，(1,0) 会被压栈两次，总和算成 5
		Name:     "square_island_has_cycle",
		Grid:     [][]int{{1, 1}, {1, 1}},
		K:        4,
		Expected: 1,
	},
	{
		// 一条从 (0,0) 出发的螺旋形路径，依次需要向右、下、左、上、右走，
		// 少写任何一个方向都会把岛屿拆成几块，每块总和都小于 14
		Name: "spiral_island_needs_all_directions",
		Grid: [][]int{
			{1, 1, 1, 1},
			{0, 0, 0, 1},
			{1, 1, 0, 1},
			{1, 0, 0, 1},
			{1, 1, 1, 1},
		},
		K:        14,
		Expected: 1,
	},
	{
		// m*n 与格子值都取上限：总和 1e11 超出 int32（截断后不能被 1e6 整除，会得到 0），
		// 递归解法的递归深度也会达到 1e5
		Name:     "full_grid_max_sum",
		Grid:     fullGrid(100, 1000, 1_000_000),
		K:        1_000_000,
		Expected: 1,
	},
}
