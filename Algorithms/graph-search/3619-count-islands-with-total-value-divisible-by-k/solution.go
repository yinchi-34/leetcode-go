package solution

// countIslands 递归 DFS：访问过的格子置 0，会修改 grid。
func countIslands(grid [][]int, k int) (ans int) {
	dirs := []struct{ x, y int }{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	m, n := len(grid), len(grid[0])

	var dfs func(int, int) int
	dfs = func(i, j int) int {
		res := grid[i][j]
		grid[i][j] = 0 // 标记 (i,j) 访问过
		for _, d := range dirs {
			x, y := i+d.x, j+d.y
			if 0 <= x && x < m && 0 <= y && y < n && grid[x][y] > 0 {
				res += dfs(x, y)
			}
		}
		return res
	}

	for i, row := range grid {
		for j, x := range row {
			if x > 0 && dfs(i, j)%k == 0 {
				ans++
			}
		}
	}
	return
}

// countIslandsStack 显式栈 DFS：用 visited 标记，入栈时标记，不修改 grid。
func countIslandsStack(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])

	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	res := 0

	type Point struct {
		r int
		c int
	}

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if grid[r][c] != 0 && !visited[r][c] {
				stack := []Point{{r, c}}
				visited[r][c] = true

				total := 0
				for len(stack) > 0 {
					point := stack[len(stack)-1]
					stack = stack[:(len(stack) - 1)]
					x, y := point.r, point.c
					total += grid[x][y]
					for _, pos := range [][2]int{
						{x + 1, y},
						{x - 1, y},
						{x, y - 1},
						{x, y + 1},
					} {
						nx, ny := pos[0], pos[1]

						if 0 <= nx && nx < m && 0 <= ny && ny < n {
							if grid[nx][ny] != 0 && !visited[nx][ny] {
								stack = append(stack, Point{nx, ny})
								visited[nx][ny] = true
							}
						}
					}
				}
				if total%k == 0 {
					res++
				}
			}
		}
	}
	return res
}
