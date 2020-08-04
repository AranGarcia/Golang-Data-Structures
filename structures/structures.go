package structures

import "fmt"

// Type definitions
type dataStructure struct {
	size   uint
	last   int
	values []float64
}

type accesor interface {
	Insert(float64)
	Pop() float64
	PrintInfo()
}

// Stack ...
type Stack struct {
	dataStructure
}

// General methods
func (ds *dataStructure) Init(size uint) {
	// Initialize structure fields
	ds.size = size
	ds.last = -1

	// Initialize the array
	ds.values = make([]float64, 0, size)
}

func (ds dataStructure) PrintValues() {
	var comma string
	fmt.Print("[")
	for i := 0; i <= ds.last; i++ {
		if i != 0 {
			comma = ", "
		} else {
			comma = ""
		}
		fmt.Printf("%v%v", comma, ds.values[i])
	}
	fmt.Print("]\n")
}

func (ds dataStructure) PrintInfo() {
	fmt.Printf("Length: %v\n", len(ds.values))
	fmt.Printf("Capacity: %v\n", cap(ds.values))
	fmt.Printf("Values: %v\n", ds.values)
}

/*
 * Use pointer receivers when a method needs to:
 * - Modify fields of a struct.
 * - When implementing polymorphism based on an interface type.
 */

/*
 * Stack methods
 */

// Insert ...
func (s *Stack) Insert(value float64) {
	s.last++
	s.values = append(s.values, value)
}

// Pop ...
func (s *Stack) Pop() float64 {
	var value = s.values[s.last]
	s.values = s.values[:s.last]
	s.last--
	return value
}
