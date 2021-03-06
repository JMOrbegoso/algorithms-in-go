package fibonacci_memoization

import (
	"bufio"
	"fmt"
	"os"

	"github.com/JMOrbegoso/algorithms-in-go/helpers"
)

const min = 0

type FibonacciSequence struct{}

func (f FibonacciSequence) Name() string {
	return "Fibonacci Sequence using Memoization"
}

func (f FibonacciSequence) Show() {
	var results []uint64

	// Request the parameters
	reader := bufio.NewReader(os.Stdin)
	max := helpers.RequestNumber(reader, "Enter the last number of the sequence of natural numbers:")

	// Get natural numbers array
	naturalNumbers := helpers.NaturalNumbersRange(min, uint16(max))

	// Use fibonacci
	for i := 0; i < len(naturalNumbers); i++ {
		result := fibonacci(uint8(i))
		results = append(results, result)
	}

	// Print the result array
	fmt.Println(results)
}
