package binary_search

import (
	"testing"
)

type testType struct {
	array []int16
	find  int16
	want  int
}

func TestBinarySearch(t *testing.T) {
	// Arrange
	testsValues := []testType{
		{[]int16{0, 1}, 1, 1},
		{[]int16{5, 10}, 1, -1},
		{[]int16{3, 6, 9, 12, 15, 28, 45}, 15, 4},
		{[]int16{-8, -1, 0, 6, 9, 12, 15, 28, 45}, 9, 4},
		{[]int16{-8, -1, 0, 6, 9, 12, 15, 28, 45}, 99, -1},
	}

	// Act
	for i := 0; i < len(testsValues); i++ {
		output := binarySearch(testsValues[i].array, 0, len(testsValues[i].array)-1, testsValues[i].find)

		// Assert
		if output != testsValues[i].want {
			t.Fatalf("on the array: %v, got %v, but it should be %v", testsValues[i].array, output, testsValues[i].want)
		}
	}
}
