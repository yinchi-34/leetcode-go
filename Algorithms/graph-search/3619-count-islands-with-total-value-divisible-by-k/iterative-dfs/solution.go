package solution

func CountIslands(grid [][]int, k int) int {
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
