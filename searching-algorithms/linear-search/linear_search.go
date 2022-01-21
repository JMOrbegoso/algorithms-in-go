package linear_search

func linearSearch(array []int16, find int16) int {
	for i := range array {
		if array[i] == find {
			return i
		}
	}

	return -1
}
