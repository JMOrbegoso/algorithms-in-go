package helpers

import (
	"bufio"
	"strconv"
	"strings"
)

func RequestNumber(reader *bufio.Reader, message string) uint64 {
	var value uint64
	requestValue := true

	for requestValue {
		requestValue = false

		println(message)

		valueString, err := reader.ReadString('\n')
		if err != nil {
			println("Invalid entered value")
			requestValue = true
		}

		if requestValue {
			continue
		}

		valueString = strings.TrimSuffix(valueString, "\n")

		value, err = strconv.ParseUint(valueString, 10, 32)
		if err != nil {
			println("Can't parse the entered number")
			requestValue = true
		}
	}

	return value
}
