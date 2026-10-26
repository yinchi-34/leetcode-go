package solution

// scoreOfParentheses uses a stack: st[k] is the score of the balanced
// blocks already closed inside the k-th open layer (st[0] is the top level).
func scoreOfParentheses(s string) int {
	st := []int{0}

	for _, c := range s {
		if c == '(' {
			st = append(st, 0)
			continue
		}
		v := st[len(st)-1]
		st = st[:len(st)-1]
		st[len(st)-1] += max(2*v, 1)
	}
	return st[0]
}

// scoreOfParenthesesDepth sums 2^d over every adjacent "()" pair,
// where d is the number of layers enclosing that pair.
func scoreOfParenthesesDepth(s string) int {
	res, depth := 0, 0

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			depth++
			continue
		}
		depth--
		if s[i-1] == '(' {
			res += 1 << depth
		}
	}
	return res
}
