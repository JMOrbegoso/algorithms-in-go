package fibonacci_memoization

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/JMOrbegoso/algorithms-in-go/helpers"
)

const min = 0

type FibonacciSequence struct{}

func (f FibonacciSequence) Name() string {
	return "Fibonacci Sequence using Memoization"
}

func (f FibonacciSequence) Show() {
	var results []uint64

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the last number of the sequence of natural numbers:")

	maxString, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Invalid entered value", err)
		return
	}
	maxString = strings.TrimSuffix(maxString, "\n")

	max, err := strconv.ParseUint(maxString, 10, 32)
	if err != nil {
		fmt.Println("Can't parse the number", err)
		return
	}

	// Get natural numbers array
	naturalNumbers := helpers.NaturalNumbersRange(min, uint16(max))

	// Declare the hash map
	cache := make(map[uint8]uint64)

	// Use fibonacci
	for i := 0; i < len(naturalNumbers); i++ {
		result := fibonacci_memoization(uint8(i), cache)
		results = append(results, result)
	}

	// Print the result array
	fmt.Println(results)
}
