package binary_search

func binarySearch(array []int16, leftLimit int, rightLimit int, find int16) int {
	if leftLimit > rightLimit {
		return -1
	}

	mid := (leftLimit + rightLimit) / 2

	if array[mid] == find {
		return mid
	}

	if find > array[mid] {
		return binarySearch(array, mid+1, rightLimit, find)
	}
	return binarySearch(array, leftLimit, mid-1, find)
}
