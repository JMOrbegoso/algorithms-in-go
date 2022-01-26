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
	for _, testValue := range testsValues {
		output := NaturalNumbersRange(testValue.min, testValue.max)

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

func TestNaturalNumbersInverseRange(t *testing.T) {
	// Arrange
	testsValues := []testType{
		{0, 2, []uint16{2, 1, 0}},
		{1, 2, []uint16{2, 1}},
		{1, 10, []uint16{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}},
		{8, 10, []uint16{10, 9, 8}},
	}

	// Act
	for _, testValue := range testsValues {
		output := NaturalNumbersInverseRange(testValue.min, testValue.max)

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
