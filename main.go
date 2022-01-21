package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	fibonacci "github.com/JMOrbegoso/algorithms-in-go/other-algorithms/fibonacci"
	fibonacci_memoization "github.com/JMOrbegoso/algorithms-in-go/other-algorithms/fibonacci-memoization"
	bubble_sort "github.com/JMOrbegoso/algorithms-in-go/sorting-algorithms/bubble-sort"
)

func main() {
	var a algorithm

	fmt.Println(strings.Repeat("-", 64))
	fmt.Println("Algorithms in Go")
	fmt.Println("Repository: https://github.com/JMOrbegoso/algorithms-in-go")
	fmt.Println("Author: JMOrbegoso (https://www.jmorbegoso.com)")
	fmt.Println(strings.Repeat("-", 64))

	fmt.Println("\nHi there 👋")
	fmt.Println("\nYou can choose any of the following algorithms:")

	fmt.Println("1. Bubble Sort")

	fmt.Println("2. Fibonacci Sequence")
	fmt.Println("3. Fibonacci Sequence using Memoization")

	fmt.Println("\n0. Close")
	fmt.Println("\nEnter your option:")

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {

		menuOption := input.Text()

		if menuOption == "0" {
			fmt.Println("\nBye bye 👋")
			return
		}

		switch menuOption {
		case "1":
			a = bubble_sort.BubbleSort{}
		case "2":
			a = fibonacci.FibonacciSequence{}
		case "3":
			a = fibonacci_memoization.FibonacciSequence{}

		default:
			fmt.Println("\nPlease, enter only any of the options that were shown before:")
		}

		showAlgorithmName(a.Name())
		a.Show()

		fmt.Println("\nEnter your option:")
	}
}

func showAlgorithmName(algorithmName string) {
	fmt.Println(strings.Repeat("-", 64))
	fmt.Println(algorithmName)
	fmt.Println(strings.Repeat("-", 64))
}
