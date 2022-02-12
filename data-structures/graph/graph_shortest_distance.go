package graph

import (
	"math"
)

type PriorityQueueItem struct {
	distance int
	vertex   *Vertex
	previous *PriorityQueueItem
}

type PriorityQueue struct {
	items []*PriorityQueueItem
	count uint
}

func (p *PriorityQueue) enqueue(current *Vertex, distance int, previous *PriorityQueueItem) {
	newItem := PriorityQueueItem{vertex: current, distance: distance, previous: previous}
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
	minDistance := math.MaxInt
	var dequeuedItem *PriorityQueueItem
	var dequeuedItemIndex uint

	for i, item := range p.items {
		if minDistance > item.distance {
			minDistance = item.distance

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

func (g *Graph) makePriorityQueueFromUnvisitedVertices(first int) *PriorityQueue {
	unvisited := &PriorityQueue{}

	for _, vertex := range g.vertices {
		if vertex.key == first {
			unvisited.enqueue(vertex, 0, nil)
		} else {
			unvisited.enqueue(vertex, math.MaxInt, nil)
		}
	}

	return unvisited
}

// Get the shortest distance between two vertices using Dijkstra’s algorithm
func (g *Graph) shortestDistance(from, to int) []*PriorityQueueItem {
	var visited []*PriorityQueueItem
	var currentVisit *PriorityQueueItem
	unvisited := g.makePriorityQueueFromUnvisitedVertices(from)

	for unvisited.count != 0 {
		currentVisit = unvisited.dequeue()

		// Visit the adjacent vertices
		for _, adjacent := range currentVisit.vertex.adjacent {
			newDistance := currentVisit.distance + adjacent.distance

			for _, priorityQueueItem := range unvisited.items {
				if priorityQueueItem.vertex == adjacent.vertex {
					if priorityQueueItem.distance > newDistance {
						priorityQueueItem.distance = newDistance
						priorityQueueItem.previous = currentVisit
					}
				}
			}
		}

		visited = append(visited, currentVisit)
	}

	return visited
}
