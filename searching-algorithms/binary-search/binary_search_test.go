package binary_search

import (
	"testing"
)

type testType struct {
	array []int16
	find  int16
	want  int
}

func TestBinarySearchOnTheEntireArray(t *testing.T) {
	// Arrange
	testsValues := []testType{
		{[]int16{0, 1}, 1, 1},
		{[]int16{5, 10}, 1, -1},
		{[]int16{3, 6, 9, 12, 15, 28, 45}, 15, 4},
		{[]int16{-8, -1, 0, 6, 9, 12, 15, 28, 45}, 9, 4},
		{[]int16{-8, -1, 0, 6, 9, 12, 15, 28, 45}, 99, -1},
	}

	// Act
	for _, testValue := range testsValues {
		output := binarySearchOnTheEntireArray(testValue.array, testValue.find)

		// Assert
		if output != testValue.want {
			t.Fatalf("on the array: %v, got %v, but it should be %v", testValue.array, output, testValue.want)
		}
	}
}
