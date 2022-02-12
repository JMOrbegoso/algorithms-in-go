package graph

import (
	"math"
	"testing"
)

func TestEnqueue(t *testing.T) {
	// Arrange
	priorityQueue := PriorityQueue{}

	// Act
	priorityQueue.enqueue(nil, 2, nil)
	priorityQueue.enqueue(nil, 1, nil)

	// Assert
	if priorityQueue.count != 2 {
		t.Errorf("expected %v, got %v", 2, priorityQueue.count)
	}

	if priorityQueue.items[0].distance != 2 {
		t.Errorf("expected %v, got %v", 2, priorityQueue.items[0].distance)
	}

	if priorityQueue.items[1].distance != 1 {
		t.Errorf("expected %v, got %v", 1, priorityQueue.items[1].distance)
	}
}

func TestSwap(t *testing.T) {
	// Arrange
	priorityQueue := PriorityQueue{}
	priorityQueue.enqueue(nil, 3, nil)
	priorityQueue.enqueue(nil, 2, nil)
	priorityQueue.enqueue(nil, 1, nil)

	// Act
	priorityQueue.swap(0, 1)

	// Assert
	if priorityQueue.count != 3 {
		t.Errorf("expected %v, got %v", 3, priorityQueue.count)
	}

	if priorityQueue.items[0].distance != 2 {
		t.Errorf("expected %v, got %v", 2, priorityQueue.items[0].distance)
	}

	if priorityQueue.items[1].distance != 3 {
		t.Errorf("expected %v, got %v", 3, priorityQueue.items[1].distance)
	}

	if priorityQueue.items[2].distance != 1 {
		t.Errorf("expected %v, got %v", 1, priorityQueue.items[2].distance)
	}
}

func TestDequeue(t *testing.T) {
	// Arrange
	priorityQueue := PriorityQueue{}
	priorityQueue.enqueue(nil, 3, nil)
	priorityQueue.enqueue(nil, 2, nil)
	priorityQueue.enqueue(nil, 1, nil)

	// Act
	output := priorityQueue.dequeue()

	// Assert
	if output.distance != 1 {
		t.Errorf("expected %v, got %v", 1, output.distance)
	}

	if priorityQueue.count != 2 {
		t.Errorf("expected %v, got %v", 2, priorityQueue.count)
	}
}

func TestMakePriorityQueueFromUnvisitedVertices(t *testing.T) {
	// Arrange
	vertex1 := Vertex{key: 10}
	vertex2 := Vertex{key: 20}
	graph := Graph{vertices: []*Vertex{&vertex1, &vertex2}}

	// Act
	priorityQueue := graph.makePriorityQueueFromUnvisitedVertices(20)

	// Assert
	if priorityQueue.items[0].vertex.key != 10 {
		t.Errorf("expected %v, got %v", 10, priorityQueue.items[0].vertex.key)
	}
	if priorityQueue.items[0].distance != math.MaxInt {
		t.Errorf("expected %v, got %v", math.MaxInt, priorityQueue.items[0].distance)
	}
	if priorityQueue.items[0].previous != nil {
		t.Errorf("expected %v, got %v", nil, priorityQueue.items[0].previous)
	}

	if priorityQueue.items[1].vertex.key != 20 {
		t.Errorf("expected %v, got %v", 20, priorityQueue.items[1].vertex.key)
	}
	if priorityQueue.items[1].distance != 0 {
		t.Errorf("expected %v, got %v", 0, priorityQueue.items[1].distance)
	}
	if priorityQueue.items[1].previous != nil {
		t.Errorf("expected %v, got %v", nil, priorityQueue.items[1].previous)
	}
}

func TestShortestDistance(t *testing.T) {
	// Arrange
	vertexS := Vertex{key: 10}
	vertexA := Vertex{key: 20}
	vertexB := Vertex{key: 30}
	vertexC := Vertex{key: 40}
	vertexD := Vertex{key: 50}
	vertexE := Vertex{key: 60}
	vertexF := Vertex{key: 70}
	vertexG := Vertex{key: 80}
	vertexH := Vertex{key: 90}
	vertexI := Vertex{key: 100}
	vertexJ := Vertex{key: 110}
	vertexK := Vertex{key: 120}
	vertexL := Vertex{key: 130}

	adjacent_S_A := Adjacent{vertex: &vertexA, distance: 7}
	adjacent_A_S := Adjacent{vertex: &vertexS, distance: 7}
	adjacent_S_B := Adjacent{vertex: &vertexB, distance: 2}
	adjacent_B_S := Adjacent{vertex: &vertexS, distance: 2}
	adjacent_A_B := Adjacent{vertex: &vertexB, distance: 3}
	adjacent_B_A := Adjacent{vertex: &vertexA, distance: 3}
	adjacent_A_D := Adjacent{vertex: &vertexD, distance: 4}
	adjacent_D_A := Adjacent{vertex: &vertexA, distance: 4}
	adjacent_B_D := Adjacent{vertex: &vertexD, distance: 4}
	adjacent_D_B := Adjacent{vertex: &vertexB, distance: 4}
	adjacent_D_F := Adjacent{vertex: &vertexF, distance: 5}
	adjacent_F_D := Adjacent{vertex: &vertexD, distance: 5}
	adjacent_H_F := Adjacent{vertex: &vertexF, distance: 3}
	adjacent_F_H := Adjacent{vertex: &vertexH, distance: 3}
	adjacent_B_H := Adjacent{vertex: &vertexH, distance: 1}
	adjacent_H_B := Adjacent{vertex: &vertexB, distance: 1}
	adjacent_H_G := Adjacent{vertex: &vertexG, distance: 2}
	adjacent_G_H := Adjacent{vertex: &vertexH, distance: 2}
	adjacent_G_E := Adjacent{vertex: &vertexE, distance: 2}
	adjacent_E_G := Adjacent{vertex: &vertexG, distance: 2}
	adjacent_S_C := Adjacent{vertex: &vertexC, distance: 3}
	adjacent_C_S := Adjacent{vertex: &vertexS, distance: 3}
	adjacent_C_L := Adjacent{vertex: &vertexL, distance: 2}
	adjacent_L_C := Adjacent{vertex: &vertexC, distance: 2}
	adjacent_L_I := Adjacent{vertex: &vertexI, distance: 4}
	adjacent_I_L := Adjacent{vertex: &vertexL, distance: 4}
	adjacent_L_J := Adjacent{vertex: &vertexJ, distance: 4}
	adjacent_J_L := Adjacent{vertex: &vertexL, distance: 4}
	adjacent_I_J := Adjacent{vertex: &vertexJ, distance: 6}
	adjacent_J_I := Adjacent{vertex: &vertexI, distance: 6}
	adjacent_I_K := Adjacent{vertex: &vertexK, distance: 4}
	adjacent_K_I := Adjacent{vertex: &vertexI, distance: 4}
	adjacent_J_K := Adjacent{vertex: &vertexK, distance: 4}
	adjacent_K_J := Adjacent{vertex: &vertexJ, distance: 4}
	adjacent_K_E := Adjacent{vertex: &vertexE, distance: 5}
	adjacent_E_K := Adjacent{vertex: &vertexK, distance: 5}

	vertexS.adjacent = []*Adjacent{&adjacent_S_A, &adjacent_S_B, &adjacent_S_C}
	vertexA.adjacent = []*Adjacent{&adjacent_A_S, &adjacent_A_B, &adjacent_A_D}
	vertexB.adjacent = []*Adjacent{&adjacent_B_S, &adjacent_B_A, &adjacent_B_D, &adjacent_B_H}
	vertexC.adjacent = []*Adjacent{&adjacent_C_S, &adjacent_C_L}
	vertexD.adjacent = []*Adjacent{&adjacent_D_A, &adjacent_D_B, &adjacent_D_F}
	vertexE.adjacent = []*Adjacent{&adjacent_E_G, &adjacent_E_K}
	vertexF.adjacent = []*Adjacent{&adjacent_F_D, &adjacent_F_H}
	vertexG.adjacent = []*Adjacent{&adjacent_G_H, &adjacent_G_E}
	vertexH.adjacent = []*Adjacent{&adjacent_H_B, &adjacent_H_F, &adjacent_H_G}
	vertexI.adjacent = []*Adjacent{&adjacent_I_L, &adjacent_I_J, &adjacent_I_K}
	vertexJ.adjacent = []*Adjacent{&adjacent_J_L, &adjacent_J_I, &adjacent_J_K}
	vertexK.adjacent = []*Adjacent{&adjacent_K_I, &adjacent_K_J, &adjacent_K_E}
	vertexL.adjacent = []*Adjacent{&adjacent_L_C, &adjacent_L_I, &adjacent_L_J}

	graph := Graph{vertices: []*Vertex{&vertexS, &vertexA, &vertexB, &vertexC, &vertexD, &vertexE, &vertexF, &vertexG, &vertexH, &vertexI, &vertexJ, &vertexK, &vertexL}}

	// Act
	output := graph.shortestDistance(10, 60)

	//Assert
	for _, travel := range output {
		if travel.vertex.key == 60 {
			if travel.distance != 7 {
				t.Errorf("expected %v, got %v", 7, travel.distance)
			}

			if travel.previous.vertex.key != 80 {
				t.Errorf("expected %v, got %v", 80, travel.previous.vertex.key)
			}
			if travel.previous.distance != 5 {
				t.Errorf("expected %v, got %v", 5, travel.previous.distance)
			}

			if travel.previous.previous.vertex.key != 90 {
				t.Errorf("expected %v, got %v", 90, travel.previous.previous.vertex.key)
			}
			if travel.previous.previous.distance != 3 {
				t.Errorf("expected %v, got %v", 3, travel.previous.previous.distance)
			}

			if travel.previous.previous.previous.vertex.key != 30 {
				t.Errorf("expected %v, got %v", 30, travel.previous.previous.previous.vertex.key)
			}
			if travel.previous.previous.previous.distance != 2 {
				t.Errorf("expected %v, got %v", 2, travel.previous.previous.previous.distance)
			}

			if travel.previous.previous.previous.previous.vertex.key != 10 {
				t.Errorf("expected %v, got %v", 10, travel.previous.previous.previous.previous.vertex.key)
			}
			if travel.previous.previous.previous.previous.distance != 0 {
				t.Errorf("expected %v, got %v", 0, travel.previous.previous.previous.previous.distance)
			}

			if travel.previous.previous.previous.previous.previous != nil {
				t.Errorf("expected %v, got %v", nil, travel.previous.previous.previous.previous.previous)
			}
		}
	}
}
