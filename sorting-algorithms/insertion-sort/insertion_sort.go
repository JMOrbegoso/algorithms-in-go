package insertion_sort

func insertionSort(array []int16) []int16 {
	var key int16

	for i := 1; i < len(array); i++ {
		key = array[i]

		for j := i - 1; 0 <= j; j-- {
			if array[j] > key {
				array[j+1] = array[j]
				array[j] = key
			}
		}
	}

	return array
}
