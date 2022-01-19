package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Repeat("-", 64))
	fmt.Println("Algorithms in Go")
	fmt.Println("Repository: https://github.com/JMOrbegoso/algorithms-in-go")
	fmt.Println("Author: JMOrbegoso (https://www.jmorbegoso.com)")
	fmt.Println(strings.Repeat("-", 64))

	fmt.Println("\nHi there 👋")
	fmt.Println("\nYou can choose any of the following algorithms:")

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
		default:
			fmt.Println("\nPlease, enter only any of the options that were shown before:")
		}

		fmt.Println("\nEnter your option:")
	}
}
