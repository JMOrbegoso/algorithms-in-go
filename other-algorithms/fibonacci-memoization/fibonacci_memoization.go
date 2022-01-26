package fibonacci_memoization

func fibonacci(i uint8) uint64 {
	cache := make(map[uint8]uint64)
	cache[0] = 0
	cache[1] = 1

	return fibonacci_memoization(i, cache)
}

func fibonacci_memoization(i uint8, cache map[uint8]uint64) uint64 {
	value, exists := cache[i]
	if exists {
		return value
	}

	result := fibonacci_memoization(i-1, cache) + fibonacci_memoization(i-2, cache)

	cache[i] = result

	return result
}
