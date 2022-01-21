package insertion_sort

import (
	"bufio"
	"fmt"
	"os"

	"github.com/JMOrbegoso/algorithms-in-go/helpers"
)

type Insertion_Sort struct{}

func (f Insertion_Sort) Name() string {
	return "Insertion Sort"
}

func (f Insertion_Sort) Show() {
	reader := bufio.NewReader(os.Stdin)

	min := helpers.RequestNumber(reader, "Enter the first number of the sequence of natural numbers:")
	max := helpers.RequestNumber(reader, "Enter the last number of the sequence of natural numbers:")

	// Build the numbers array
	naturalNumbers := helpers.NaturalNumbersInverseRange(uint16(min), uint16(max))

	// Convert all the numbers of the array to int16
	numbersRange := make([]int16, len(naturalNumbers))
	for i := range numbersRange {
		numbersRange[i] = int16(naturalNumbers[i])
	}
	fmt.Println("The original array was", numbersRange)

	// Use insertion sort
	result := insertionSort(numbersRange)

	// Print the result array
	fmt.Println("The sorted array is", result)
}
