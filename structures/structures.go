/*
Package structures provides implementations of basic data structures.
- Stack: First in, last out
*/
package structures

import (
	"fmt"
	"sort"
)

// DataStructure is the base type for structures that will need to implement
// how the data is organized and accesed.
type DataStructure struct {
	values []float64
}

// Interface type for all data structures that will implement retrieval
// methods for data access.
type dataAccesor interface {
	Pop() float64
	Peek() float64
	Sort()
	PrintInfo()
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

// Init pre-allocates the memory used for the data structure
func (ds *DataStructure) Init(capacity uint) {
	ds.values = make([]float64, 0, capacity)
}

// InitWithValues initializes the data structure with an existing slice of values.
func (ds *DataStructure) InitWithValues(values []float64) {
	ds.values = make([]float64, len(values))
	copy(ds.values, values)
}

// GetSize returns the amount of values appended into the structure.
func (ds *DataStructure) GetSize() int {
	return int(len(ds.values))
}

// Insert adds a value into the structure.
func (ds *DataStructure) Insert(value float64) {
	ds.values = append(ds.values, value)
}

// PrintInfo displays information about the data structure.
func (ds *DataStructure) PrintInfo() {
	fmt.Printf("Length: %v\n", len(ds.values))
	fmt.Printf("Capacity: %v\n", cap(ds.values))
	fmt.Printf("Values: %v\n", ds.values)
}

// PrintValues prints the values in a comma separated list.
func (ds DataStructure) PrintValues() {
	var comma string
	fmt.Print("[")
	for i := 0; i < len(ds.values); i++ {
		if i != 0 {
			comma = ", "
		} else {
			comma = ""
		}
		fmt.Printf("%v%v", comma, ds.values[i])
	}
	fmt.Print("]\n")
}

// Sort reorganizes the values by sorting them in ascending order.
func (ds *DataStructure) Sort() {
	sort.Float64s(ds.values)
}

/*
 * Use pointer receivers when a method needs to:
 * - Modify fields of a struct.
 * - When implementing polymorphism based on an interface type.
 */

/*
 * Stack methods
 */

// Peek returns the next value to be returned without removing it
// from the stack (that is, the most recent value inserted).
func (s *Stack) Peek() float64 {
	var last int = len(s.values) - 1
	return s.values[last]
}

// Pop returns and removes the last value introduced into stack.
func (s *Stack) Pop() float64 {
	var last int = len(s.values) - 1
	var value float64 = s.values[last]
	s.values = s.values[:last]
	return value
}

/*
 * Queue methods
 */

// Peek returns the next value to be returned without removing it
// from the queue (that is, the first value inserted).
func (q *Queue) Peek() float64 {
	return q.values[0]
}

// Pop implementation of the queue returns the first item that was
// introduced in the queue.
func (q *Queue) Pop() float64 {
	var value float64 = q.values[0]
	q.values = q.values[1:]
	return value
}
