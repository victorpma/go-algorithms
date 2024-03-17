package binarytree

type BinaryTree struct {
	root *Node
}

type Node struct {
	value int
	left  *Node
	right *Node
}

// NewNode creates a Node with empty left and right Nodes.
func NewNode(value int) *Node {
	return &Node{
		value: value,
		left:  nil,
		right: nil,
	}
}

// Insert adds a new Node in the tree
func (b *BinaryTree) Insert(value int) {
	if b.root == nil {
		b.root = NewNode(value)
	} else {
		b.root.insert(value)
	}
}

func (n *Node) insert(value int) {
	if value < n.value {
		if n.left == nil {
			n.left = NewNode(value)
		} else {
			n.left.insert(value)
		}
	} else {
		if n.right == nil {
			n.right = NewNode(value)
		} else {
			n.right.insert(value)
		}
	}
}

// Search searchs a node by value
func (b *BinaryTree) Search(target int) *Node {
	return b.search(b.root, target)
}

func (b *BinaryTree) search(curretNode *Node, target int) *Node {
	if curretNode.value == target {
		return curretNode
	}

	if target < curretNode.value && curretNode.left != nil {
		return b.search(curretNode.left, target)
	} else if curretNode.right != nil {
		return b.search(curretNode.right, target)
	}

	return nil
}

// Value returns the node value
func (n *Node) Value() int {
	return n.value
}
