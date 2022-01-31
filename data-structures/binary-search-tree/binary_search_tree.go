package binary_search_tree

type Node struct {
	value int
	left  *Node
	right *Node
}

func (n *Node) include(value int) bool {
	if n == nil {
		return false
	}

	if n.value == value {
		return true
	}

	if n.value > value {
		return n.left.include(value)
	}
	if value > n.value {
		return n.right.include(value)
	}

	return false
}

func (n *Node) insert(value int) {

	if n == nil {
		return
	}

	if n.value > value {
		if n.left == nil {
			n.left = &Node{value: value}
		} else {
			n.left.insert(value)
		}
	}
	if value > n.value {
		if n.right == nil {
			n.right = &Node{value: value}
		} else {
			n.right.insert(value)
		}
	}
}
