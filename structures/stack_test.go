package structures

import "testing"

func TestStackWithValues(t *testing.T) {
	var initialValues []float64 = []float64{1, 1, 2, 3, 5, 8, 13}
	var stack Stack = Stack{}

	stack.InitWithValues(initialValues)

	if stack.Size() != len(initialValues) {
		t.Errorf("stack.InitWithValues failed: expected %v values but got %v\n", len(initialValues), stack.Size())
	} else {
		t.Logf("stack.InitWithValues succesfully initialized with %v values.\n", stack.Size())
	}
}

func TestPeekAndPop(t *testing.T) {
	stack := &Stack{}

	stack.Insert(2)
	stack.Insert(4)
	stack.Insert(6)
	stack.Insert(8)
	stack.Insert(10)

	if stack.Peek() != 10 {
		t.Errorf("stack.Peek failed: expected 10 but got %v\n", stack.Peek())
	}

	stack.Pop()              // Removes 10
	stack.Pop()              // Removes 8
	lastValue := stack.Pop() // Removes 6

	if lastValue != 6 {
		t.Errorf("stack.Peek failed: expected 6 but got %v\n", lastValue)
	}

	if stack.Peek() != 4 {
		t.Errorf("stack.Peek failed: expected 4 but got %v\n", stack.Peek())
	}

	stack.Pop() // Removes 4
	stack.Pop() //Removes 2

	// Test empty stack
	if stack.Size() != 0 {
		t.Errorf("stack.Size: expected to be empty but got size %v instead", stack.Size())
	}
}

func TestStackSlice(t *testing.T) {
	var stack *Stack = new(Stack)

	var values []float64 = []float64{10, 2, 7, 3.1416, -8}

	for _, val := range values {
		stack.Insert(val)
	}

	// Compare the slice output from the stack
	var stackSlice []float64 = stack.Slice()
	for i, stackValue := range stackSlice {
		if values[i] != stackValue {
			t.Errorf("stack.Slice failed: found difference at index %v.\n", i)
			return
		}
	}
	t.Logf("stack.Sice succesfully created slice: %v: \n", stackSlice)
}
