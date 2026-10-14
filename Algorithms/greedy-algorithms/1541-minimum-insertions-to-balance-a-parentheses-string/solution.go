package main 

func minAddToMakeValid(s string) int {
	res, right := 0, 0

	for _, c := range s {
		if c == '(' {
			if right % 2 == 1 {
				res++
				right--
			}
			right += 2
		}
		if c == ')' {
			right--
			if right < 0 {
				res++
				right += 2
			}
		} 
	}
	return res + right
}