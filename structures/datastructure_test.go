package structures

import "testing"

func TestEmptyInit(t *testing.T) {
	var ds DataStructure = DataStructure{}

	// Test insert on empty structure
	var size int = ds.Size()
	if size != 0 {
		t.Errorf("Data structure not initialized properly: size %v", size)
	}

	var dsSlice []float64 = ds.Slice()
	size = len(dsSlice)
	if size != 0 {
		t.Errorf("Slice generated from empty data structure is not empty.")
	}
}

func TestInitWithValues(t *testing.T) {
	var initialValues []float64 = []float64{1, 1, 2, 3, 5, 8, 13, 21}
	var ds *DataStructure = new(DataStructure)

	// Test that the data structure can initialize
	ds.InitWithValues(initialValues)
	var size int = ds.Size()
	if size != len(initialValues) {
		t.Errorf("Error initalizing the size of the data structure: expected %v but got %v\n",
			len(initialValues),
			ds.Size(),
		)
	}

	var dsValues []float64 = ds.Slice()
	size = len(dsValues)

	// Test size of the slice generated from the data structure
	if size != len(initialValues) {
		t.Errorf("DataStructure.Slice returns a different size: expected %v but got %v.\n",
			len(initialValues),
			size,
		)
	}

	// Test slice genarated from data structure is the same as the initial values.
	for i, value := range dsValues {
		if initialValues[i] != value {
			t.Errorf("Slice at index %v should be %v but has value %v.\n",
				i, initialValues[i], value,
			)
		}
	}
}

func TestInsert(t *testing.T) {
	var ds DataStructure = DataStructure{}

	ds.Insert(-20)
	ds.Insert(10)
	ds.Insert(0)

	var size int = ds.Size()
	if size != 3 {
		t.Errorf("DataStructure length must be 3 but got size %v.", size)
	}

	ds.Insert(10)
	ds.Insert(50)
	ds.Insert(100)

	size = ds.Size()
	if size != 6 {
		t.Errorf("DataStructure length must be 6 but got size %v.", size)
	}
}
