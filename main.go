package main

import (
	"fmt"

	"github.com/AranGarcia/golang-data-structures/structures"
)

func testStack() {
	fmt.Println("Testing the stack data structure...")
	var stack *structures.Stack = new(structures.Stack)
	stack.Init(10)

	stack.Insert(1)
	stack.PrintValues()
	stack.Insert(3)
	stack.Insert(4)
	fmt.Printf("Next value to be popped: %v\n", stack.Peek())
	value := stack.Pop()

	fmt.Printf("Popped: %v\n", value)

	stack.PrintValues()

	stack.PrintInfo()
}

func testQueue() {
	fmt.Println("Testing the queue data structure...")
	var values = []float64{1, 1, 2, 3, 5, 8}
	var queue *structures.Queue = new(structures.Queue)

	queue.InitWithValues(values)
	queue.PrintInfo()

	fmt.Printf("Next value to be popped: %v\n", queue.Peek())
	fmt.Printf("Popped: %v\n", queue.Pop())
	queue.Pop()
	queue.Pop()
	queue.Pop()
	queue.Pop()
	fmt.Printf("Next value to be popped: %v\n", queue.Peek())
	fmt.Printf("Popped: %v\n", queue.Pop())
}

func main() {
	testStack()
	fmt.Println("---------------------")
	testQueue()
}
