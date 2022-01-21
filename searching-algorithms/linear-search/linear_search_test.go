package linear_search

import (
	"testing"
)

type testType struct {
	array []int16
	find  int16
	want  int
}

func TestLinearSearch(t *testing.T) {
	// Arrange
	testsValues := []testType{
		{[]int16{0, 1}, 1, 1},
		{[]int16{5, 10}, 1, -1},
		{[]int16{3, 8, 3, 3, 5, 8, 5}, 3, 0},
		{[]int16{30, 18, 10, 20, 3}, 3, 4},
		{[]int16{30, 18, 10, 20, 3}, 1, -1},
	}

	// Act
	for i := 0; i < len(testsValues); i++ {
		output := linearSearch(testsValues[i].array, testsValues[i].find)

		// Assert
		if output != testsValues[i].want {
			t.Fatalf("on the array: %v, got %v, but it should be %v", testsValues[i].array, output, testsValues[i].want)
		}
	}
}
