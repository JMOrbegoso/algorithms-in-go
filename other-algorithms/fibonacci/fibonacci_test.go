package fibonacci

import "testing"

type testType struct {
	input uint8
	want  uint64
}

func TestFibonacci(t *testing.T) {
	// Arrange
	testsValues := []testType{
		{0, 0},
		{1, 1},
		{2, 1},
		{10, 55},
		{15, 610},
		{30, 832040},
		{40, 102334155},
	}

	// Act
	for _, testValue := range testsValues {
		output := fibonacci(testValue.input)

		// Assert
		if output != testValue.want {
			t.Fatalf("when the input is %v, it should return %v, but got %v", testValue.input, testValue.want, output)
		}
	}
}
