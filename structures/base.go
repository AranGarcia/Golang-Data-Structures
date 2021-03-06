package structures

import (
	"fmt"
)

/*
 * Use pointer receivers when a method needs to:
 * - Modify fields of a struct.
 * - When implementing polymorphism based on an interface type.
 */

// Node is the basic unit for the data structure. Nodes will be connected to
// by pointers to other nodes and each data structure will handle theses links.
type Node struct {
	value    float64
	previous *Node
	next     *Node
}

// DataStructure is the base type for structures that will need to implement
// how the data is organized and accesed.
type DataStructure struct {
	size  int
	first *Node
	last  *Node
}

// DataAccesor provides basic interface to shared methods among data structures.
type DataAccesor interface {
	Clear()
	InitWithValues([]float64)
	Insert(float64)
	PrintInfo()
	PrintValues()
	Size() int
	Slice() []float64
}

/*
 * DataStructure methods
 */

// Clear deletes all nodes in the datastructure.
func (ds *DataStructure) Clear() {
	ds.size = 0
	ds.first = nil
	ds.last = nil
}

// InitWithValues initializes the data structure with an existing slice of values.
func (ds *DataStructure) InitWithValues(values []float64) {
	if len(values) == 0 {
		return
	}
	var node *Node = &Node{
		value:    values[0],
		previous: nil,
		next:     nil,
	}

	ds.first = node
	ds.size = 1

	var temp *Node
	for i := 1; i < len(values); i++ {
		temp = &Node{
			value:    values[i],
			previous: node,
			next:     nil,
		}
		node.next = temp
		node = temp
		ds.size++
	}

	ds.last = node
}

// Insert adds a value into the structure.
func (ds *DataStructure) Insert(newValue float64) {
	var aux *Node = &Node{
		value:    newValue,
		previous: ds.last,
		next:     nil,
	}

	if ds.last != nil {
		// Re-arrange existing nodes
		ds.last.next = aux
	} else {
		// stack is empty
		ds.first = aux
	}
	ds.last = aux
	ds.size++
}

// PrintInfo displays information about the data structure.
func (ds DataStructure) PrintInfo() {
	fmt.Printf("Length: %v\n", ds.size)
}

// PrintValues prints the values in a comma separated list.
func (ds DataStructure) PrintValues() {
	var comma string
	fmt.Print("[")

	var aux *Node = ds.first
	for aux != nil {
		if aux.next == nil {
			comma = ""
		} else {
			comma = ", "
		}
		fmt.Printf("%v%v", aux.value, comma)
		aux = aux.next
	}
	fmt.Print("]\n")
}

// Size returns the amount of nodes in the data structure.
func (ds DataStructure) Size() int {
	return ds.size
}

// Slice converts the data structure into a slice
func (ds DataStructure) Slice() []float64 {
	var values []float64 = make([]float64, ds.size)

	var index int = 0
	var aux *Node = ds.first
	for aux != nil {
		values[index] = aux.value
		aux = aux.next
		index++
	}

	return values
}
