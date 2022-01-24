package factorial

func factorial(i uint16) uint32 {
	if i == 0 {
		return 1
	}

	return uint32(i) * factorial(i-1)
}
