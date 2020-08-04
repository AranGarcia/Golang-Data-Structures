/*
Package structures provides implementations of basic data structures.
- Stack: First in, last out
*/
package structures

import "fmt"

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
	PrintInfo()
}

// Stack is the data structure that provides a last in, first out
// (LIFO) data organization.
type Stack struct {
	DataStructure
}

// Init pre-allocates the memory used for the data structure
func (ds *DataStructure) Init(capacity uint) {
	ds.values = make([]float64, 0, capacity)
}

// PrintValues ...
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

// PrintInfo displays information about the data structure.
func (ds *DataStructure) PrintInfo() {
	fmt.Printf("Length: %v\n", len(ds.values))
	fmt.Printf("Capacity: %v\n", cap(ds.values))
	fmt.Printf("Values: %v\n", ds.values)
}

// Insert adds a value into the structure.
func (ds *DataStructure) Insert(value float64) {
	ds.values = append(ds.values, value)
}

// GetSize returns the amount of values appended into the structure.
func (ds *DataStructure) GetSize() int {
	return int(len(ds.values))
}

/*
 * Use pointer receivers when a method needs to:
 * - Modify fields of a struct.
 * - When implementing polymorphism based on an interface type.
 */

/*
 * Stack methods
 */

// Pop returns and removes the last value introduced into stack.
func (s *Stack) Pop() float64 {
	var last int = len(s.values) - 1
	var value float64 = s.values[last]
	s.values = s.values[:last]
	return value
}
