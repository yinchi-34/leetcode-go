package solution

import "strings"

func createGrid(m int, n int) []string {
	res := make([]string, m)
	res[0] = strings.Repeat(".", n)
	for i := 1; i < m; i++ {
		res[i] = strings.Repeat("#", n-1) + "."
	}
	return res
}
