package quick_sort

func quickSortTheEntireArray(array []int16) []int16 {
	quickSort(array, 0, len(array)-1)
	return array
}

func quickSort(array []int16, leftLimit int, rightLimit int) {
	if leftLimit >= rightLimit {
		return
	}

	pivotIndex := partition(array, leftLimit, rightLimit)

	quickSort(array, leftLimit, pivotIndex-1)
	quickSort(array, pivotIndex+1, rightLimit)
}

func partition(array []int16, leftLimit int, rightLimit int) int {
	pivotIndex := rightLimit
	i := leftLimit - 1

	for j := leftLimit; j < rightLimit; j++ {
		if array[j] <= array[pivotIndex] {
			i++
			swap(&array[i], &array[j])
		}
	}

	i++
	swap(&array[i], &array[pivotIndex])

	return i
}

func swap(a *int16, b *int16) {
	temp := *a
	*a = *b
	*b = temp
}
