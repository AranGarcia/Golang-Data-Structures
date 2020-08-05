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

// Stack is the data structure that provides a last in, first out
// (LIFO) data organization.
type Stack struct {
	DataStructure
}

// Queue is the data structure that provides a first in, first out
// (FIFO) data organization.
type Queue struct {
	DataStructure
}

/*
 * DataStructure methods
 */

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
func (ds *DataStructure) Insert(value float64) {
	//ds.values = append(ds.values, value)
	var aux *Node = ds.last
	ds.last = &Node{
		value:    value,
		previous: aux,
		next:     nil,
	}

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

// Sort reorganizes the values by sorting them in ascending order.
func (ds *DataStructure) Sort() {
	fmt.Println("Not yet implemented.")
}
