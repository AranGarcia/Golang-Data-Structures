package structures

/*
 * Queue methods
 */

// Peek returns the next value to be returned without removing it
// from the queue (that is, the first value inserted).
func (q *Queue) Peek() float64 {
	return q.first.value
}

// Pop implementation of the queue returns the first item that was
// introduced in the queue.
func (q *Queue) Pop() float64 {
	var value float64 = q.first.value
	q.first = q.first.next
	q.size--
	return value
}
