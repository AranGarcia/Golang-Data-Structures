package main

import (
	"fmt"

	"github.com/AranGarcia/golang-data-structures/structures"
)

func stackExample() {
	fmt.Println("Testing the stack data structure...")
	var stack *structures.Stack = new(structures.Stack)

	var initialVales []float64 = []float64{1, 1, 2, 3, 5, 8}

	stack.InitWithValues(initialVales)
	stack.PrintInfo()
	stack.PrintValues()
}

func main() {
	stackExample()
}
