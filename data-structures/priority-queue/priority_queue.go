package priority_queue

import (
	"math"
)

type PriorityQueueItem struct {
	value    int
	priority int
}

type PriorityQueue struct {
	items []*PriorityQueueItem
	count uint
}

func (p *PriorityQueue) enqueue(value, priority int) {
	newItem := PriorityQueueItem{value: value, priority: priority}
	p.items = append(p.items, &newItem)
	p.count++
}

func (p *PriorityQueue) swap(i, j uint) {
	if i >= p.count || j >= p.count {
		return
	}

	item1 := &p.items[i]
	item2 := &p.items[j]

	temp := *item1
	*item1 = *item2
	*item2 = temp
}

func (p *PriorityQueue) dequeue() *PriorityQueueItem {
	maxPriority := math.MinInt
	var dequeuedItem *PriorityQueueItem
	var dequeuedItemIndex uint

	for i, item := range p.items {
		if item.priority > maxPriority {
			maxPriority = item.priority

			dequeuedItem = item
			dequeuedItemIndex = uint(i)
		}
	}

	// Remove the dequeued element from the array
	p.swap(0, dequeuedItemIndex)
	p.items = p.items[1:]
	p.count--

	return dequeuedItem
}
