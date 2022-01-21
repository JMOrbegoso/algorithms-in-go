package bubble_sort

import (
	"testing"
)

type testType struct {
	input []int16
	want  []int16
}

func TestSwap(t *testing.T) {
	// Arrange
	inputA := int16(1)
	inputB := int16(10)

	// Act
	swap(&inputA, &inputB)

	// Assert
	if inputA != 10 {
		t.Fatalf("got %v, it should be %v", inputA, 10)
	}

	if inputB != 1 {
		t.Fatalf("got %v, it should be %v", inputB, 1)
	}
}

func TestBubbleSort(t *testing.T) {
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
	for i := 0; i < len(testsValues); i++ {
		output := bubbleSort(testsValues[i].input)

		// Assert
		if len(output) != len(testsValues[i].want) {
			t.Fatalf("the output length is %v, it should be %v", len(output), len(testsValues[i].want))
		}

		for j := 0; j < len(testsValues[i].want); j++ {
			if output[j] != testsValues[i].want[j] {
				t.Fatalf("got %v, it should be %v", output[j], testsValues[i].want[j])
			}
		}
	}
}
