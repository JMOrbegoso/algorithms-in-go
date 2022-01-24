package factorial

import "testing"

type testType struct {
	input uint16
	want  uint32
}

func TestFactorial(t *testing.T) {
	// Arrange
	testsValues := []testType{
		{1, 1},
		{2, 2},
		{5, 120},
		{6, 720},
		{10, 3628800},
	}

	// Act
	for i := 0; i < len(testsValues); i++ {
		output := factorial(testsValues[i].input)

		// Assert
		if output != testsValues[i].want {
			t.Fatalf("when the input is %v, it should return %v, but got %v", testsValues[i].input, testsValues[i].want, output)
		}
	}
}
