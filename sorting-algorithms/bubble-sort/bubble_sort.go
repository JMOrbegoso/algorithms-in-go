package bubble_sort

func bubbleSort(array []int16) []int16 {
	swapRequired := true

	for swapRequired {
		swapRequired = false

		for i := 0; i < len(array)-1; i++ {
			if array[i] > array[i+1] {
				swapRequired = true
				swap(&array[i], &array[i+1])
			}
		}
	}

	return array
}

func swap(a *int16, b *int16) {
	temp := *a
	*a = *b
	*b = temp
}
