package invert_string

import "testing"

type testType struct {
	input string
	want  string
}

func TestInvert(t *testing.T) {
	// Arrange
	testsValues := []testType{
		{"a", "a"},
		{"ab", "ba"},
		{"abc", "cba"},
		{"abcd", "dcba"},
		{"456", "654"},
		{"4567", "7654"},
	}

	// Act
	for _, testValue := range testsValues {
		output := invert(testValue.input)

		// Assert
		if output != testValue.want {
			t.Fatalf("when the input is %v, it should return %v, but got %v", testValue.input, testValue.want, output)
		}
	}
}
