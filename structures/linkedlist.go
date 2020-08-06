package structures

// LinkedList is a double-linked data structure that can retrieve values
// from the start or the end.
type LinkedList struct {
	DataStructure
}

// Get returns the value given the index
func (ll *LinkedList) Get(index int) float64 {
	if index > ll.size {
		// Out of index: raise error
	}
	var node *Node = ll.first
	for i := 0; i < index; i++ {
		node = node.next
	}

	return node.value
}

// RGet (for lack of a better name) returns a value given
// an index starting from the last node
func (ll *LinkedList) RGet(index int) float64 {
	if index > ll.size {
		// Out of index: raise error
	}

	var node *Node = ll.last
	for i := 0; i < index; i++ {
		node = node.previous
	}

	return node.value
}

// Pop returns a value from the structure given its index and removes it.
func (ll *LinkedList) Pop(index int) float64 {
	if index > ll.size {
		// Out of index: raise error
	}
	var node *Node = ll.first
	for i := 0; i < index; i++ {
		node = node.next
	}

	node.previous.next = node.next.previous

	return node.value
}
