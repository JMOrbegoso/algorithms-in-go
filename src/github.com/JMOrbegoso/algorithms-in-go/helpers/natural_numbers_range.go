package helpers

// Return an array of natural numbers from the minimum to the maximum number entered as parameters
func NaturalNumbersRange(min uint16, max uint16) []uint16 {
	var numbersRange []uint16

	for i := min; i <= max; i++ {
		numbersRange = append(numbersRange, i)
	}

	return numbersRange
}
