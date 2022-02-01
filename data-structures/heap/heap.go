package heap

type MaxHeap struct {
	items []int
}

func getParentIndex(childIndex int) int {
	return (childIndex - 1) / 2
}

func getLeftChildIndex(parentIndex int) int {
	return parentIndex*2 + 1
}

func getRightChildIndex(parentIndex int) int {
	return parentIndex*2 + 2
}

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func (h *MaxHeap) insert(value int) {
	h.items = append(h.items, value)
	h.heapifyBottomToTop(len(h.items) - 1)
}

func (h *MaxHeap) heapifyBottomToTop(childIndex int) {
	for h.items[childIndex] > h.items[getParentIndex(childIndex)] {
		parentIndex := getParentIndex(childIndex)

		swap(&h.items[childIndex], &h.items[parentIndex])
		childIndex = parentIndex
	}
}

func (h *MaxHeap) extract() int {
	if len(h.items) == 0 {
		return -1
	}

	extracted := h.items[0]

	lastIndex := len(h.items) - 1

	swap(&h.items[0], &h.items[lastIndex])

	h.items = h.items[:lastIndex]

	h.heapifyTopToBottom()

	return extracted
}

func (h *MaxHeap) heapifyTopToBottom() {
	parentIndex := 0
	lastIndex := len(h.items) - 1
	leftChildIndex := getLeftChildIndex(parentIndex)
	rightChildIndex := getRightChildIndex(parentIndex)
	greaterChildIndex := 0

	for leftChildIndex <= lastIndex {
		if leftChildIndex == lastIndex {
			greaterChildIndex = leftChildIndex
		} else if h.items[leftChildIndex] > h.items[rightChildIndex] {
			greaterChildIndex = leftChildIndex
		} else {
			greaterChildIndex = rightChildIndex
		}

		// Check if the greater child is greater than its parent
		if h.items[greaterChildIndex] > h.items[parentIndex] {
			swap(&h.items[greaterChildIndex], &h.items[parentIndex])

			parentIndex = greaterChildIndex
			leftChildIndex = getLeftChildIndex(parentIndex)
			rightChildIndex = getRightChildIndex(parentIndex)
		} else {
			return
		}
	}
}
