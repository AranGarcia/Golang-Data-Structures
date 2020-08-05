package structures

/*
 * Stack methods
 */

// Peek returns the next value to be returned without removing it
// from the stack (that is, the most recent value inserted).
func (s *Stack) Peek() float64 {
	var last int = len(s.values) - 1
	return s.values[last]
}

// Pop returns and removes the last value introduced into stack.
func (s *Stack) Pop() float64 {
	var last int = len(s.values) - 1
	var value float64 = s.values[last]
	s.values = s.values[:last]
	return value
}
