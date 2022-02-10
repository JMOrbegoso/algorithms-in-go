package graph

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestVertexWasVisited(t *testing.T) {
	// Arrange
	vertex2 := Vertex{key: 20}
	vertex1 := Vertex{key: 10}
	visitedVertices := make([]*Vertex, 2)
	visitedVertices = append(visitedVertices, &vertex1)

	// Act
	output1 := vertexWasVisited(visitedVertices, &vertex1)
	output2 := vertexWasVisited(visitedVertices, &vertex2)

	// Assert
	if output1 != true {
		t.Errorf("expected %v, got %v", true, output1)
	}
	if output2 != false {
		t.Errorf("expected %v, got %v", false, output2)
	}
}

func TestPrint(t *testing.T) {
	// Arrange
	vertex6 := Vertex{key: 60}
	vertex5 := Vertex{key: 50}
	adjacent1 := Adjacent{vertex: &vertex6, distance: 1}
	vertex4 := Vertex{key: 40, adjacent: []*Adjacent{&adjacent1}}
	adjacent2 := Adjacent{vertex: &vertex5, distance: 1}
	vertex3 := Vertex{key: 30, adjacent: []*Adjacent{&adjacent2}}
	adjacent3 := Adjacent{vertex: &vertex4, distance: 1}
	vertex2 := Vertex{key: 20, adjacent: []*Adjacent{&adjacent3}}
	adjacent4 := Adjacent{vertex: &vertex2, distance: 1}
	adjacent5 := Adjacent{vertex: &vertex3, distance: 1}
	vertex1 := Vertex{key: 10, adjacent: []*Adjacent{&adjacent4, &adjacent5}}
	graph := Graph{vertices: []*Vertex{&vertex1, &vertex2, &vertex3, &vertex4, &vertex5, &vertex6}}

	r, w, _ := os.Pipe()
	os.Stdout = w

	// Act
	graph.print()

	w.Close()
	out, _ := ioutil.ReadAll(r)

	// Assert
	if !strings.Contains(string(out), strconv.Itoa(vertex1.key)) {
		t.Errorf("expected %v", vertex1.key)
	}
	if !strings.Contains(string(out), strconv.Itoa(vertex2.key)) {
		t.Errorf("expected %v", vertex2.key)
	}
	if !strings.Contains(string(out), strconv.Itoa(vertex3.key)) {
		t.Errorf("expected %v", vertex3.key)
	}
	if !strings.Contains(string(out), strconv.Itoa(vertex4.key)) {
		t.Errorf("expected %v", vertex4.key)
	}
	if !strings.Contains(string(out), strconv.Itoa(vertex5.key)) {
		t.Errorf("expected %v", vertex5.key)
	}
	if !strings.Contains(string(out), strconv.Itoa(vertex6.key)) {
		t.Errorf("expected %v", vertex6.key)
	}
}

func TestPrintCyclicGraph(t *testing.T) {
	// Arrange
	vertex1 := Vertex{key: 10}
	vertex2 := Vertex{key: 20}
	vertex3 := Vertex{key: 30}

	adjacent1 := Adjacent{vertex: &vertex2, distance: 1}
	adjacent2 := Adjacent{vertex: &vertex3, distance: 1}
	adjacent3 := Adjacent{vertex: &vertex1, distance: 1}
	adjacent4 := Adjacent{vertex: &vertex3, distance: 1}
	adjacent5 := Adjacent{vertex: &vertex1, distance: 1}
	adjacent6 := Adjacent{vertex: &vertex2, distance: 1}

	vertex1.adjacent = []*Adjacent{&adjacent1, &adjacent2}
	vertex2.adjacent = []*Adjacent{&adjacent3, &adjacent4}
	vertex3.adjacent = []*Adjacent{&adjacent5, &adjacent6}

	graph := Graph{vertices: []*Vertex{&vertex1, &vertex2, &vertex3}}

	r, w, _ := os.Pipe()
	os.Stdout = w

	// Act
	graph.print()

	w.Close()
	out, _ := ioutil.ReadAll(r)

	// Assert
	if !strings.Contains(string(out), strconv.Itoa(vertex1.key)) {
		t.Errorf("expected %v", vertex1.key)
	}
	if !strings.Contains(string(out), strconv.Itoa(vertex2.key)) {
		t.Errorf("expected %v", vertex2.key)
	}
	if !strings.Contains(string(out), strconv.Itoa(vertex3.key)) {
		t.Errorf("expected %v", vertex3.key)
	}
}

func TestAddVertex(t *testing.T) {
	// Arrange
	graph := Graph{}

	// Act
	graph.addVertex(10)
	graph.addVertex(20)

	// Assert
	if graph.vertices[0].key != 10 {
		t.Errorf("expected %v, got %v", 10, graph.vertices[0].key)
	}
	if graph.vertices[1].key != 20 {
		t.Errorf("expected %v, got %v", 20, graph.vertices[1].key)
	}
}

func TestGetVertexNotFound(t *testing.T) {
	// Arrange
	vertex1 := Vertex{key: 10}
	graph := Graph{vertices: []*Vertex{&vertex1}}

	// Act
	output := graph.getVertex(100)

	// Assert
	if output != nil {
		t.Errorf("expected %v, got %v", nil, output)
	}
}

func TestGetVertex(t *testing.T) {
	// Arrange
	vertex1 := Vertex{key: 10}
	graph := Graph{vertices: []*Vertex{&vertex1}}

	// Act
	output := graph.getVertex(10)

	// Assert
	if output.key != 10 {
		t.Errorf("expected %v, got %v", 10, output.key)
	}
}

func TestAddDuplicatedEdge(t *testing.T) {
	// Arrange
	vertex2 := Vertex{key: 20}
	adjacent1 := Adjacent{vertex: &vertex2, distance: 1}
	vertex1 := Vertex{key: 10, adjacent: []*Adjacent{&adjacent1}}
	graph := Graph{vertices: []*Vertex{&vertex1, &vertex2}}

	// Act
	graph.addEdge(10, 20, 1)

	// Assert
	if len(vertex1.adjacent) != 1 {
		t.Errorf("expected %v, got %v", 1, len(vertex1.adjacent))
	}
	if vertex1.adjacent[0].vertex.key != vertex2.key {
		t.Errorf("expected %v, got %v", vertex2.key, vertex1.adjacent[0].vertex.key)
	}

}

func TestAddEdge(t *testing.T) {
	// Arrange
	vertex2 := Vertex{key: 20}
	vertex1 := Vertex{key: 10}
	graph := Graph{vertices: []*Vertex{&vertex1, &vertex2}}

	// Act
	graph.addEdge(10, 20, 18)

	// Assert
	if len(vertex1.adjacent) != 1 {
		t.Errorf("expected %v, got %v", 1, len(vertex1.adjacent))
	}
	if vertex1.adjacent[0].vertex.key != vertex2.key {
		t.Errorf("expected %v, got %v", vertex2.key, vertex1.adjacent[0].vertex.key)
	}
	if vertex1.adjacent[0].distance != 18 {
		t.Errorf("expected %v, got %v", 18, vertex1.adjacent[0].distance)
	}
}

func TestContains(t *testing.T) {
	// Arrange
	vertex2 := Vertex{key: 20}
	vertex1 := Vertex{key: 10}
	graph := Graph{vertices: []*Vertex{&vertex1, &vertex2}}

	// Act
	output := graph.contains(10)

	// Assert
	if output != true {
		t.Errorf("expected %v, got %v", true, output)
	}
}

func TestContainsNotFound(t *testing.T) {
	// Arrange
	vertex2 := Vertex{key: 20}
	vertex1 := Vertex{key: 10}
	graph := Graph{vertices: []*Vertex{&vertex1, &vertex2}}

	// Act
	output := graph.contains(1000)

	// Assert
	if output != false {
		t.Errorf("expected %v, got %v", false, output)
	}
}

func TestHasPath(t *testing.T) {
	// Arrange
	vertex6 := Vertex{key: 60}
	adjacent1 := Adjacent{vertex: &vertex6, distance: 1}
	vertex5 := Vertex{key: 50, adjacent: []*Adjacent{&adjacent1}}
	adjacent2 := Adjacent{vertex: &vertex5, distance: 1}
	vertex4 := Vertex{key: 40, adjacent: []*Adjacent{&adjacent2}}
	adjacent3 := Adjacent{vertex: &vertex4, distance: 1}
	vertex3 := Vertex{key: 30, adjacent: []*Adjacent{&adjacent3}}
	adjacent4 := Adjacent{vertex: &vertex3, distance: 1}
	vertex2 := Vertex{key: 20, adjacent: []*Adjacent{&adjacent4}}
	adjacent5 := Adjacent{vertex: &vertex2, distance: 1}
	vertex1 := Vertex{key: 10, adjacent: []*Adjacent{&adjacent5}}
	graph := Graph{vertices: []*Vertex{&vertex1, &vertex2, &vertex3, &vertex4, &vertex5, &vertex6}}

	// Act
	output := graph.hasPath(20, 60)

	// Assert
	if output != true {
		t.Errorf("expected %v, got %v", true, output)
	}
}

func TestHasNoPath(t *testing.T) {
	// Arrange
	vertex7 := Vertex{key: 70}
	vertex6 := Vertex{key: 60}
	adjacent1 := Adjacent{vertex: &vertex6, distance: 1}
	vertex5 := Vertex{key: 50, adjacent: []*Adjacent{&adjacent1}}
	adjacent2 := Adjacent{vertex: &vertex5, distance: 1}
	vertex4 := Vertex{key: 40, adjacent: []*Adjacent{&adjacent2}}
	adjacent3 := Adjacent{vertex: &vertex4, distance: 1}
	vertex3 := Vertex{key: 30, adjacent: []*Adjacent{&adjacent3}}
	adjacent4 := Adjacent{vertex: &vertex3, distance: 1}
	vertex2 := Vertex{key: 20, adjacent: []*Adjacent{&adjacent4}}
	adjacent5 := Adjacent{vertex: &vertex2, distance: 1}
	vertex1 := Vertex{key: 10, adjacent: []*Adjacent{&adjacent5}}
	graph := Graph{vertices: []*Vertex{&vertex1, &vertex2, &vertex3, &vertex4, &vertex5, &vertex6, &vertex7}}

	// Act
	output := graph.hasPath(20, 70)

	// Assert
	if output != false {
		t.Errorf("expected %v, got %v", false, output)
	}
}

func TestVisitAdjacentVertices(t *testing.T) {
	// Arrange
	vertex3 := Vertex{key: 30}
	adjacent1 := Adjacent{vertex: &vertex3, distance: 1}
	vertex2 := Vertex{key: 20, adjacent: []*Adjacent{&adjacent1}}
	adjacent2 := Adjacent{vertex: &vertex2, distance: 1}
	vertex1 := Vertex{key: 10, adjacent: []*Adjacent{&adjacent2}}
	var visitedVertices []*Vertex

	// Act
	visitAdjacentVertices(&visitedVertices, &vertex1)

	// Assert
	if visitedVertices[0] != &vertex1 {
		t.Errorf("expected %v, got %v", &vertex1, visitedVertices[0])
	}
	if visitedVertices[1] != &vertex2 {
		t.Errorf("expected %v, got %v", &vertex2, visitedVertices[1])
	}
	if visitedVertices[2] != &vertex3 {
		t.Errorf("expected %v, got %v", &vertex3, visitedVertices[2])
	}
}

func TestConnectedComponentsCount(t *testing.T) {
	// Arrange
	vertex7 := Vertex{key: 70}
	vertex6 := Vertex{key: 60}
	adjacent1 := Adjacent{vertex: &vertex6, distance: 1}
	vertex5 := Vertex{key: 50, adjacent: []*Adjacent{&adjacent1}}
	adjacent2 := Adjacent{vertex: &vertex5, distance: 1}
	vertex4 := Vertex{key: 40, adjacent: []*Adjacent{&adjacent2}}
	vertex3 := Vertex{key: 30}
	adjacent3 := Adjacent{vertex: &vertex3, distance: 1}
	vertex2 := Vertex{key: 20, adjacent: []*Adjacent{&adjacent3}}
	adjacent4 := Adjacent{vertex: &vertex2, distance: 1}
	vertex1 := Vertex{key: 10, adjacent: []*Adjacent{&adjacent4}}
	graph := Graph{vertices: []*Vertex{&vertex1, &vertex2, &vertex3, &vertex4, &vertex5, &vertex6, &vertex7}}

	// Act
	output := graph.connectedComponentsCount()

	// Assert
	if output != 3 {
		t.Errorf("expected %v, got %v", 3, output)
	}
}

func TestShortestPath(t *testing.T) {
	// Arrange
	vertex4 := Vertex{key: 40}
	adjacent1 := Adjacent{vertex: &vertex4, distance: 1}
	vertex5 := Vertex{key: 50, adjacent: []*Adjacent{&adjacent1}}
	adjacent2 := Adjacent{vertex: &vertex4, distance: 1}
	vertex3 := Vertex{key: 30, adjacent: []*Adjacent{&adjacent2}}
	adjacent3 := Adjacent{vertex: &vertex3, distance: 1}
	vertex2 := Vertex{key: 20, adjacent: []*Adjacent{&adjacent3}}
	adjacent4 := Adjacent{vertex: &vertex2, distance: 1}
	adjacent5 := Adjacent{vertex: &vertex5, distance: 1}
	vertex1 := Vertex{key: 10, adjacent: []*Adjacent{&adjacent4, &adjacent5}}
	graph := Graph{vertices: []*Vertex{&vertex1, &vertex2, &vertex3, &vertex4, &vertex5}}

	// Act
	output := graph.shortestPath(10, 40)

	//Assert
	if output != 2 {
		t.Errorf("expected %v, got %v", 2, output)
	}
}

func TestShortestPathNotFound(t *testing.T) {
	// Arrange
	vertex6 := Vertex{key: 60}
	vertex4 := Vertex{key: 40}
	adjacent1 := Adjacent{vertex: &vertex4, distance: 1}
	vertex5 := Vertex{key: 50, adjacent: []*Adjacent{&adjacent1}}
	adjacent2 := Adjacent{vertex: &vertex4, distance: 1}
	vertex3 := Vertex{key: 30, adjacent: []*Adjacent{&adjacent2}}
	adjacent3 := Adjacent{vertex: &vertex3, distance: 1}
	vertex2 := Vertex{key: 20, adjacent: []*Adjacent{&adjacent3}}
	adjacent4 := Adjacent{vertex: &vertex2, distance: 1}
	adjacent5 := Adjacent{vertex: &vertex5, distance: 1}
	vertex1 := Vertex{key: 10, adjacent: []*Adjacent{&adjacent4, &adjacent5}}
	graph := Graph{vertices: []*Vertex{&vertex1, &vertex2, &vertex3, &vertex4, &vertex5, &vertex6}}

	// Act
	output := graph.shortestPath(10, 60)

	//Assert
	if output != -1 {
		t.Errorf("expected %v, got %v", -1, output)
	}
}
