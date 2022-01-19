package fibonacci

func fibonacci(i uint8) uint64 {
	if i == 0 {
		return 0
	}

	if i == 1 {
		return 1
	}

	return fibonacci(i-1) + fibonacci(i-2)
}
