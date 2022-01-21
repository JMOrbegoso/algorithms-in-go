package selection_sort

func selectionSort(array []int16) []int16 {
	for i := range array {
		min := &array[i]

		for j := i + 1; j < len(array); j++ {
			if *min > array[j] {
				min = &array[j]
			}
		}

		swap(min, &array[i])
	}

	return array
}

func swap(a *int16, b *int16) {
	temp := *a
	*a = *b
	*b = temp
}
