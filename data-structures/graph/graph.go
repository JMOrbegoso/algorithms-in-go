package graph

import "fmt"

type Adjacent struct {
	vertex   *Vertex
	distance int
}

type Vertex struct {
	key      int
	adjacent []*Adjacent
}

type Graph struct {
	vertices []*Vertex
}

type queueNode struct {
	vertex   *Vertex
	next     *queueNode
	distance int
}

type queue struct {
	front  *queueNode
	rear   *queueNode
	length uint
}

func (q *queue) enqueue(node *queueNode) {
	node.next = nil

	if q.rear == nil {
		q.front = node
	} else {
		q.rear.next = node
	}

	q.rear = node
	q.length++
}

func (q *queue) dequeue() *queueNode {
	if q.front == nil {
		return nil
	}

	dequeuedNode := q.front

	if q.front == q.rear {
		q.front = nil
		q.rear = nil
	} else {
		q.front = q.front.next
	}

	q.length--

	return dequeuedNode
}

func vertexWasVisited(visitedVertices []*Vertex, vertex *Vertex) bool {
	for _, visited := range visitedVertices {
		if visited == vertex {
			return true
		}
	}
	return false
}

// Print the graph using BFS
func (g *Graph) print() {
	if len(g.vertices) == 0 {
		return
	}

	visitedVertices := make([]*Vertex, len(g.vertices))

	auxQueue := queue{}
	var dequeuedVertex *Vertex
	var newQueueNode *queueNode

	newQueueNode = &queueNode{vertex: g.vertices[0]}
	auxQueue.enqueue(newQueueNode)

	for auxQueue.length != 0 {
		dequeuedVertex = auxQueue.dequeue().vertex
		visitedVertices = append(visitedVertices, dequeuedVertex)

		fmt.Println(dequeuedVertex.key)

		for _, adjacent := range dequeuedVertex.adjacent {
			if vertexWasVisited(visitedVertices, adjacent.vertex) {
				continue
			}

			visitedVertices = append(visitedVertices, adjacent.vertex)

			newQueueNode = &queueNode{vertex: adjacent.vertex}
			auxQueue.enqueue(newQueueNode)
		}
	}
}

func (g *Graph) addVertex(key int) {
	newVertex := &Vertex{key: key}
	g.vertices = append(g.vertices, newVertex)
}

func (g *Graph) getVertex(key int) *Vertex {
	for _, vertex := range g.vertices {
		if vertex.key == key {
			return vertex
		}
	}
	return nil
}

func (g *Graph) addEdge(from, to, distance int) {
	vertexFrom := g.getVertex(from)
	if vertexFrom == nil {
		return
	}

	vertexTo := g.getVertex(to)
	if vertexTo == nil {
		return
	}

	// Validate if vertexFrom already have an edge to vertexTo
	for _, adjacent := range vertexFrom.adjacent {
		if adjacent.vertex == vertexTo {
			return
		}
	}

	// Add the new edge
	newAdjacent := Adjacent{vertex: vertexTo, distance: distance}
	vertexFrom.adjacent = append(vertexFrom.adjacent, &newAdjacent)
}

func (g *Graph) contains(key int) bool {
	for _, vertex := range g.vertices {
		if vertex.key == key {
			return true
		}
	}

	return false
}

func (g *Graph) hasPath(from, to int) bool {
	vertexFrom := g.getVertex(from)
	if vertexFrom == nil {
		return false
	}

	vertexTo := g.getVertex(to)
	if vertexTo == nil {
		return false
	}

	var dequeuedVertex *Vertex
	var newQueueNode *queueNode

	auxQueue := queue{}

	newQueueNode = &queueNode{vertex: vertexFrom}
	auxQueue.enqueue(newQueueNode)

	for auxQueue.length != 0 {
		dequeuedVertex = auxQueue.dequeue().vertex

		if dequeuedVertex == vertexTo {
			return true
		}

		for _, adjacent := range dequeuedVertex.adjacent {
			newQueueNode = &queueNode{vertex: adjacent.vertex}
			auxQueue.enqueue(newQueueNode)
		}
	}

	return false
}

func visitAdjacentVertices(visitedVertices *[]*Vertex, vertex *Vertex) {
	if vertexWasVisited(*visitedVertices, vertex) {
		return
	}

	*visitedVertices = append(*visitedVertices, vertex)

	for _, adjacent := range vertex.adjacent {
		visitAdjacentVertices(visitedVertices, adjacent.vertex)
	}
}

func (g *Graph) connectedComponentsCount() int {
	count := 0

	if len(g.vertices) == 0 {
		return count
	}

	visitedVertices := make([]*Vertex, len(g.vertices))

	for _, vertex := range g.vertices {
		if vertexWasVisited(visitedVertices, vertex) {
			continue
		}

		count++

		visitAdjacentVertices(&visitedVertices, vertex)
	}

	return count
}

func (g *Graph) shortestPath(from, to int) int {
	vertexFrom := g.getVertex(from)
	if vertexFrom == nil {
		return -1
	}

	vertexTo := g.getVertex(to)
	if vertexTo == nil {
		return -1
	}

	var dequeuedNode *queueNode
	var newQueueNode *queueNode

	auxQueue := queue{}

	newQueueNode = &queueNode{vertex: vertexFrom}
	auxQueue.enqueue(newQueueNode)

	for auxQueue.length != 0 {
		dequeuedNode = auxQueue.dequeue()

		if dequeuedNode.vertex == vertexTo {
			return dequeuedNode.distance
		}

		for _, adjacent := range dequeuedNode.vertex.adjacent {
			newQueueNode = &queueNode{vertex: adjacent.vertex, distance: dequeuedNode.distance + 1}
			auxQueue.enqueue(newQueueNode)
		}
	}

	return -1
}
