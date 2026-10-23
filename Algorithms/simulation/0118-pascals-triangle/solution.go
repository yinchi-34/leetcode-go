package solution

func generate(numRows int) [][]int {
	res := make([][]int, numRows)

	for i := range res {
		res[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				res[i][j] = 1
			} else {
				res[i][j] = res[i-1][j] + res[i-1][j-1]
			}
		}
	}
	return res
}
