package main

import (
	"fmt"

	"github.com/AranGarcia/golang-data-structures/structures"
)

func stackExample() {
	fmt.Println("Testing the stack data structure...")
	var stack1 *structures.Stack = new(structures.Stack)

	var initialVales []float64 = []float64{1, 1, 2, 3, 5, 8}

	stack1.InitWithValues(initialVales)
	stack1.Insert(13)
	stack1.PrintInfo()
	stack1.PrintValues()
	fmt.Println(stack1.Slice())
}

func main() {
	stackExample()
}
