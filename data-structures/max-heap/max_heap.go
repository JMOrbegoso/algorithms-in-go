package max_heap

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

func (h *MaxHeap) swap(i, j int) {
	if i < 0 || j < 0 {
		return
	}

	if i >= len(h.items) || j >= len(h.items) {
		return
	}

	item1 := &h.items[i]
	item2 := &h.items[j]

	temp := *item1
	*item1 = *item2
	*item2 = temp
}

func (h *MaxHeap) insert(value int) {
	h.items = append(h.items, value)
	h.heapifyBottomToTop()
}

func (h *MaxHeap) heapifyBottomToTop() {
	childIndex := len(h.items) - 1
	parentIndex := getParentIndex(childIndex)

	for h.items[parentIndex] < h.items[childIndex] {
		h.swap(parentIndex, childIndex)

		childIndex = parentIndex
		parentIndex = getParentIndex(childIndex)
	}
}

func (h *MaxHeap) extract() int {
	if len(h.items) == 0 {
		return -1
	}

	extracted := h.items[0]

	lastIndex := len(h.items) - 1

	h.swap(0, lastIndex)

	h.items = h.items[:lastIndex]

	h.heapifyTopToBottom()

	return extracted
}

func (h *MaxHeap) heapifyTopToBottom() {
	parentIndex := 0
	leftChildIndex := getLeftChildIndex(parentIndex)
	rightChildIndex := getRightChildIndex(parentIndex)
	var greaterChildIndex int
	lastIndex := len(h.items) - 1

	for lastIndex >= leftChildIndex {
		if leftChildIndex == lastIndex {
			greaterChildIndex = leftChildIndex
		} else if h.items[leftChildIndex] > h.items[rightChildIndex] {
			greaterChildIndex = leftChildIndex
		} else {
			greaterChildIndex = rightChildIndex
		}

		// Check if the greater child is greater than its parent
		if h.items[parentIndex] < h.items[greaterChildIndex] {
			h.swap(parentIndex, greaterChildIndex)

			parentIndex = greaterChildIndex
			leftChildIndex = getLeftChildIndex(parentIndex)
			rightChildIndex = getRightChildIndex(parentIndex)
		} else {
			return
		}
	}
}
