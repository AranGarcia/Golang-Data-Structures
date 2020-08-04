package main

import (
	"fmt"

	"github.com/AranGarcia/golang-data-structures/structures"
)

func main() {
	fmt.Println("Testing data structures...")
	var stack *structures.Stack = new(structures.Stack)
	stack.Init(10)

	stack.Insert(1)
	stack.PrintValues()
	stack.Insert(3)
	stack.Insert(4)
	value := stack.Pop()

	fmt.Printf("Popped: %v\n", value)

	stack.PrintValues()

	stack.PrintInfo()
}
