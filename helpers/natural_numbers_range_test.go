package helpers

import (
	"testing"
)

type testType struct {
	min  uint16
	max  uint16
	want []uint16
}

func TestNaturalNumbersRange(t *testing.T) {
	// Arrange
	testsValues := []testType{
		{0, 2, []uint16{0, 1, 2}},
		{1, 2, []uint16{1, 2}},
		{1, 10, []uint16{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{8, 10, []uint16{8, 9, 10}},
	}

	// Act
	for i := 0; i < len(testsValues); i++ {
		output := NaturalNumbersRange(testsValues[i].min, testsValues[i].max)

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
