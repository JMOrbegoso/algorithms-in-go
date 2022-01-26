package merge_sort

import (
	"testing"
)

type testType struct {
	input []int16
	want  []int16
}

func TestMergeSortTheEntireArray(t *testing.T) {
	// Arrange
	testsValues := []testType{
		{[]int16{0, 1}, []int16{0, 1}},
		{[]int16{5, -5}, []int16{-5, 5}},
		{[]int16{-8, -3, -5}, []int16{-8, -5, -3}},
		{[]int16{3, 8, 3, 3, 5, 8, 5}, []int16{3, 3, 3, 5, 5, 8, 8}},
		{[]int16{30, 18, -10, 20, 3}, []int16{-10, 3, 18, 20, 30}},
		{[]int16{5, 4, 3, 2, 1, 0, -1, -2, -3}, []int16{-3, -2, -1, 0, 1, 2, 3, 4, 5}},
	}

	// Act
	for _, testValue := range testsValues {
		output := mergeSortTheEntireArray(testValue.input)

		// Assert
		if len(output) != len(testValue.want) {
			t.Fatalf("the output length is %v, it should be %v", len(output), len(testValue.want))
		}

		for j := 0; j < len(testValue.want); j++ {
			if output[j] != testValue.want[j] {
				t.Fatalf("got %v, it should be %v", output[j], testValue.want[j])
			}
		}
	}
}
