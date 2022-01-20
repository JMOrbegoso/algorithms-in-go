package fibonacci_memoization

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
		{50, 12586269025},
		{100, 3736710778780434371},
	}

	// Act
	for i := 0; i < len(testsValues); i++ {
		cache := make(map[uint8]uint64)

		output := fibonacci_memoization(testsValues[i].input, cache)

		// Assert
		if output != testsValues[i].want {
			t.Fatalf("when the input is %v, it should return %v, but got %v", testsValues[i].input, testsValues[i].want, output)
		}
	}
}
