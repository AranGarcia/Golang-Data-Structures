package structures

import "testing"

func TestQueuePeekAndPop(t *testing.T) {
	stack := &Queue{}

	stack.Insert(2)
	stack.Insert(4)
	stack.Insert(6)
	stack.Insert(8)
	stack.Insert(10)

	if stack.Peek() != 2 {
		t.Errorf("stack.Peek failed: expected 10 but got %v\n", stack.Peek())
	}

	stack.Pop()              // Removes 2
	stack.Pop()              // Removes 4
	lastValue := stack.Pop() // Removes 6

	if lastValue != 6 {
		t.Errorf("stack.Peek failed: expected 6 but got %v\n", lastValue)
	}

	if stack.Peek() != 8 {
		t.Errorf("stack.Peek failed: expected 4 but got %v\n", stack.Peek())
	}

	stack.Pop() // Removes 8
	stack.Pop() // Removes 10

	// Test empty stack
	if stack.Size() != 0 {
		t.Errorf("stack.Size: expected to be empty but got size %v instead", stack.Size())
	}
}
