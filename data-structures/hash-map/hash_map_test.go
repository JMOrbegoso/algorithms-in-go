package hash_map

import "testing"

func TestHashMapNotInitialized(t *testing.T) {
	// Arrange

	// Act
	hashMap := HashMap{}

	// Assert
	if hashMap.size != 0 {
		t.Errorf("expected %v, got %v", 0, hashMap.size)
	}
	if hashMap.buckets != nil {
		t.Errorf("expected %v, got %v", nil, hashMap.buckets)
	}
}

func TestInit(t *testing.T) {
	// Arrange
	hashMap := HashMap{}

	// Act
	hashMap.init(3)

	// Assert
	if hashMap.size != 3 {
		t.Errorf("expected %v, got %v", 3, hashMap.size)
	}
	if hashMap.buckets == nil {
		t.Errorf("expected not nil, got %v", hashMap.buckets)
	}
}

func TestPushNodeOnBucket(t *testing.T) {
	// Arrange
	newBucket := bucket{}
	node1 := bucketNode{key: "Jhon", value: 32}
	node2 := bucketNode{key: "Doe", value: 46}

	// Act
	newBucket.push(&node1)
	newBucket.push(&node2)

	// Assert
	if newBucket.count != 2 {
		t.Errorf("expected %v, got %v", 2, newBucket.count)
	}
	if newBucket.head.key == node1.key {
		t.Errorf("expected %v, got %v", node1.key, newBucket.head.key)
	}
	if newBucket.head.value == node1.value {
		t.Errorf("expected %v, got %v", node1.value, newBucket.head.value)
	}
	if newBucket.head.next.key == node2.key {
		t.Errorf("expected %v, got %v", node2.key, newBucket.head.next.key)
	}
	if newBucket.head.next.value == node2.value {
		t.Errorf("expected %v, got %v", node2.value, newBucket.head.next.value)
	}
	if newBucket.head.next.next != nil {
		t.Errorf("expected %v, got %v", nil, newBucket.head.next.next)
	}
}

func TestDeleteBucketNodeNotFound(t *testing.T) {
	// Arrange
	newBucket := bucket{}
	node1 := bucketNode{key: "Jhon", value: 32}
	newBucket.push(&node1)

	// Assert
	newBucket.delete("Doe")

	// Act
	if newBucket.count != 1 {
		t.Errorf("expected %v, got %v", 1, newBucket.count)
	}
	if newBucket.head.key != node1.key {
		t.Errorf("expected %v, got %v", node1.key, newBucket.head.key)
	}
	if newBucket.head.next != nil {
		t.Errorf("expected %v, got %v", nil, newBucket.head.next)
	}
}

func TestDeleteBucketHeadNode(t *testing.T) {
	// Arrange
	newBucket := bucket{}
	node1 := bucketNode{key: "Jhon", value: 32}
	node2 := bucketNode{key: "Doe", value: 46}
	node3 := bucketNode{key: "Jhon Doe", value: 55}
	newBucket.push(&node3)
	newBucket.push(&node2)
	newBucket.push(&node1)

	// Assert
	newBucket.delete(node1.key)

	// Act
	if newBucket.count != 2 {
		t.Errorf("expected %v, got %v", 2, newBucket.count)
	}
	if newBucket.head.key != node2.key {
		t.Errorf("expected %v, got %v", node2.key, newBucket.head.key)
	}
	if newBucket.head.next.key != node3.key {
		t.Errorf("expected %v, got %v", node3.key, newBucket.head.next.key)
	}
	if newBucket.head.next.next != nil {
		t.Errorf("expected %v, got %v", nil, newBucket.head.next.next)
	}
}

func TestDeleteBucketMiddleNode(t *testing.T) {
	// Arrange
	newBucket := bucket{}
	node1 := bucketNode{key: "Jhon", value: 32}
	node2 := bucketNode{key: "Doe", value: 46}
	node3 := bucketNode{key: "Jhon Doe", value: 55}
	newBucket.push(&node3)
	newBucket.push(&node2)
	newBucket.push(&node1)

	// Assert
	newBucket.delete(node2.key)

	// Act
	if newBucket.count != 2 {
		t.Errorf("expected %v, got %v", 2, newBucket.count)
	}
	if newBucket.head.key != node1.key {
		t.Errorf("expected %v, got %v", node1.key, newBucket.head.key)
	}
	if newBucket.head.next.key != node3.key {
		t.Errorf("expected %v, got %v", node3.key, newBucket.head.next.key)
	}
	if newBucket.head.next.next != nil {
		t.Errorf("expected %v, got %v", nil, newBucket.head.next.next)
	}
}

func TestDeleteBucketTailNode(t *testing.T) {
	// Arrange
	newBucket := bucket{}
	node1 := bucketNode{key: "Jhon", value: 32}
	node2 := bucketNode{key: "Doe", value: 46}
	node3 := bucketNode{key: "Jhon Doe", value: 55}
	newBucket.push(&node3)
	newBucket.push(&node2)
	newBucket.push(&node1)

	// Assert
	newBucket.delete(node3.key)

	// Act
	if newBucket.count != 2 {
		t.Errorf("expected %v, got %v", 2, newBucket.count)
	}
	if newBucket.head.key != node1.key {
		t.Errorf("expected %v, got %v", node1.key, newBucket.head.key)
	}
	if newBucket.head.next.key != node2.key {
		t.Errorf("expected %v, got %v", node2.key, newBucket.head.next.key)
	}
	if newBucket.head.next.next != nil {
		t.Errorf("expected %v, got %v", nil, newBucket.head.next.next)
	}
}

func TestHash(t *testing.T) {
	// Arrange
	hashMap := HashMap{}
	hashMap.init(3)

	// Act
	output := hashMap.hash("Doe")

	// Assert
	if output != 1 {
		t.Errorf("expected %v, got %v", 1, output)
	}
}

func TestInsert(t *testing.T) {
	// Arrange
	hashMap := HashMap{}
	hashMap.init(3)

	// Act
	hashMap.insert("Jhon", 32)
	hashMap.insert("Doe", 36)

	// Assert
	if hashMap.buckets[0].head.key != "Jhon" {
		t.Errorf("expected %v, got %v", "Jhon", hashMap.buckets[0].head.key)
	}
	if hashMap.buckets[0].head.value != 32 {
		t.Errorf("expected %v, got %v", 32, hashMap.buckets[0].head.value)
	}
	if hashMap.buckets[1].head.key != "Doe" {
		t.Errorf("expected %v, got %v", "Doe", hashMap.buckets[1].head.key)
	}
	if hashMap.buckets[1].head.value != 36 {
		t.Errorf("expected %v, got %v", 36, hashMap.buckets[1].head.value)
	}
}

func TestExistNotFound(t *testing.T) {
	// Arrange
	hashMap := HashMap{}
	hashMap.init(3)
	hashMap.insert("Jhon", 32)
	hashMap.insert("Doe", 36)

	// Act
	output := hashMap.exist("Jhon Doe")

	// Assert
	if output != false {
		t.Errorf("expected %v, got %v", false, output)
	}
}

func TestExist(t *testing.T) {
	// Arrange
	hashMap := HashMap{}
	hashMap.init(3)
	hashMap.insert("Jhon", 32)
	hashMap.insert("Doe", 36)

	// Act
	output := hashMap.exist("Jhon")

	// Assert
	if output != true {
		t.Errorf("expected %v, got %v", true, output)
	}
}

func TestDelete(t *testing.T) {
	// Arrange
	hashMap := HashMap{}
	hashMap.init(2)
	hashMap.insert("Jhon", 32)
	hashMap.insert("Doe", 36)

	// Act
	hashMap.delete("Jhon")
	hashMap.delete("Doe")

	// Assert
	if hashMap.buckets[0].head != nil {
		t.Errorf("expected %v, got %v", nil, hashMap.buckets[0].head)
	}
	if hashMap.buckets[1].head != nil {
		t.Errorf("expected %v, got %v", nil, hashMap.buckets[1].head)
	}
}
