package solution

func scoreofParentheses(s string) int {
	// Initialize a stack to keep track of the scores
	stack := []int{0} // Start with a base score of 0

	for _, c := range s {
		if c == '(' {
			// Push a new score onto the stack for the new pair of parentheses
			stack = append(stack, 0)
		} else {
			// Pop the top score from the stack
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack[len(stack)-1] += max(2*top, 1) // Update the score for the enclosing parentheses
		}
	}
	return stack[0] // The final score is at the bottom of the stack
}
