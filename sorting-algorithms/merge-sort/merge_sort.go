package merge_sort

func mergeSortTheEntireArray(array []int16) []int16 {
	return mergeSort(array, 0, len(array)-1)
}

func mergeSort(array []int16, leftLimit int, rightLimit int) []int16 {
	sort(array, leftLimit, rightLimit)
	return array
}

func sort(array []int16, leftLimit int, rightLimit int) {
	if leftLimit >= rightLimit {
		return
	}

	mid := (leftLimit + rightLimit) / 2

	mergeSort(array, leftLimit, mid)
	mergeSort(array, mid+1, rightLimit)

	merge(array, leftLimit, mid, rightLimit)
}

func merge(array []int16, leftLimit int, mid int, rightLimit int) []int16 {
	leftSubarray := createLeftSubArray(array, leftLimit, mid)
	rightSubarray := createRightSubArray(array, mid, rightLimit)

	mergeSubArraysIntoArray(array, leftSubarray, rightSubarray, leftLimit, rightLimit)

	return array
}

func createLeftSubArray(array []int16, leftLimit int, mid int) []int16 {
	leftSubarray := make([]int16, mid-leftLimit+1)

	// Copy the values
	for i := range leftSubarray {
		leftSubarray[i] = array[leftLimit+i]
	}

	return leftSubarray
}

func createRightSubArray(array []int16, mid int, rightLimit int) []int16 {
	rightSubarray := make([]int16, rightLimit-mid)

	// Copy the values
	for i := range rightSubarray {
		rightSubarray[i] = array[mid+i+1]
	}

	return rightSubarray
}

func mergeSubArraysIntoArray(array []int16, subArray1 []int16, subArray2 []int16, leftLimit int, rightLimit int) {
	i := 0
	j := 0
	k := leftLimit

	for i < len(subArray1) && j < len(subArray2) {
		if subArray1[i] <= subArray2[j] {
			array[k] = subArray1[i]
			i++
		} else {
			array[k] = subArray2[j]
			j++
		}
		k++
	}

	// Copy the remaining elements of subArray1 if there is still any
	for i < len(subArray1) {
		array[k] = subArray1[i]
		i++
		k++
	}

	// Copy the remaining elements of subArray2 if there is still any
	for j < len(subArray2) {
		array[k] = subArray2[j]
		j++
		k++
	}
}
