package helpers

// Return an array of natural numbers from the minimum to the maximum number entered as parameters
func NaturalNumbersRange(min uint16, max uint16) []uint16 {
	numbersRange := make([]uint16, max-min+1)

	for i := range numbersRange {
		numbersRange[i] = min + uint16(i)
	}

	return numbersRange
}

// Return the array of natural numbers from the maximum to the minimum number entered as parameters
func NaturalNumbersInverseRange(min uint16, max uint16) []uint16 {
	numbersRange := make([]uint16, max-min+1)

	for i := range numbersRange {
		numbersRange[i] = max - uint16(i)
	}

	return numbersRange
}
