package hash_map

type bucketNode struct {
	key   string
	value int
	next  *bucketNode
}

type bucket struct {
	head  *bucketNode
	count int
}

type HashMap struct {
	size    int
	buckets []*bucket
}

func (h *HashMap) init(size int) {
	h.size = size
	h.buckets = make([]*bucket, h.size)

	for i := range h.buckets {
		h.buckets[i] = &bucket{}
	}
}

func (b *bucket) push(node *bucketNode) {
	node.next = b.head
	b.head = node
	b.count++
}

func (b *bucket) delete(key string) {
	if b.head == nil {
		return
	}

	if b.head.key == key {
		b.head = b.head.next
		b.count--
		return
	}

	node := b.head
	for node.next != nil {
		if node.next.key == key {
			node.next = node.next.next
			b.count--
			return
		}
		node = node.next
	}
}

func (h *HashMap) hash(key string) int {
	sum := 0

	for _, v := range key {
		sum += int(v)
	}

	return sum % h.size
}

func (h *HashMap) insert(key string, value int) {
	index := h.hash(key)
	newNode := bucketNode{key: key, value: value, next: nil}
	h.buckets[index].push(&newNode)
}

func (h *HashMap) exist(key string) bool {
	index := h.hash(key)
	bucket := h.buckets[index]

	node := bucket.head
	for node != nil {
		if node.key == key {
			return true
		}
		node = node.next
	}

	return false
}

func (h *HashMap) delete(key string) {
	index := h.hash(key)
	h.buckets[index].delete(key)
}
