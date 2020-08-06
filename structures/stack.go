package structures

// Stack is the data structure that provides a last in, first out
// (LIFO) data organization.
type Stack struct {
	DataStructure
}

/*
 * Stack methods
 */

// Peek returns the next value to be returned without removing it
// from the stack (that is, the most recent value inserted).
func (s *Stack) Peek() float64 {
	return s.last.value
}

// Pop returns and removes the last value introduced into stack.
func (s *Stack) Pop() float64 {
	var value float64 = s.last.value
	s.last = s.last.previous
	s.size--
	return value
}
