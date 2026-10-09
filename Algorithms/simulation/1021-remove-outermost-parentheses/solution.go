package main

func removeOuterParentheses(S string) string {
	res := []byte{}
	depth := 0
	for _, c := range S {
		if c == '(' {
			if depth > 0 {
				res = append(res, byte(c))
			}
			depth++
		} else {
			depth --
			if depth > 0 {
				res = append(res, byte(c))
			}
		}
	}
	return string(res)
}
