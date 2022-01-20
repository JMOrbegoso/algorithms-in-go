package fibonacci_memoization

func fibonacci_memoization(i uint8, cache map[uint8]uint64) uint64 {
	if i == 0 {
		return 0
	}

	if i == 1 {
		return 1
	}

	value, exists := cache[i]
	if exists {
		return value
	}

	result := fibonacci_memoization(i-1, cache) + fibonacci_memoization(i-2, cache)

	cache[i] = result

	return result
}
