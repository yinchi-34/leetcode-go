package solution

func generate(numRows int) [][]int {
	ret := make([][]int, numRows)

	for i := range ret {
		ret[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				ret[i][j] = 1
			} else {
				ret[i][j] = ret[i-1][j] + ret[i-1][j-1]
			}
		}
	}
	return ret
}
