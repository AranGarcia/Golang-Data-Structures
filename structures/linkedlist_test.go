package structures

import "testing"

func TestLinkedList(t *testing.T) {
	ll := new(LinkedList)
	ll.Insert(1)
	ll.Insert(2)
	ll.Insert(3)
	ll.Insert(5)
	ll.Insert(10)
	ll.Insert(20)

	value := ll.Pop(4)
	if value != 10 {
		t.Errorf("ll.pop: expected 20 but got %v\n", value)
	}

	value = ll.Pop(2)
	if value != 3 {
		t.Errorf("ll.pop: expected 3 but got %v\n", value)
	}
}
