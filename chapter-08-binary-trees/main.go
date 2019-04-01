package main

import "fmt"

func main() {
	t := NewTree()
	r := NewNode(50, 1.5)
	t.root = r
	t.insert(10, 1.5)
	t.insert(20, 1.5)
	t.insert(60, 1.5)
	t.inOrder(r)
}

type Node struct {
	iData                 int
	fData                 float64
	leftChild, rightChild *Node
}

func NewNode(id int, dd float64) *Node {
	return &Node{
		iData:      id,
		fData:      dd,
		leftChild:  nil,
		rightChild: nil,
	}
}

func (n *Node) display() {
	fmt.Printf("< Node - idata: %d, fdata: %f>\n", n.iData, n.fData)
}

type Tree struct {
	root *Node
}

func NewTree() *Tree {
	return &Tree{
		root: nil,
	}
}

func (t *Tree) find(key int) *Node {
	current := t.root

	for current.iData != key {
		if current == nil {
			return nil
		}
		if key < current.iData {
			current = current.leftChild
		} else {
			current = current.rightChild
		}
	}

	return current
}

func (t *Tree) insert(id int, dd float64) {
	nn := NewNode(id, dd)

	if t.root == nil {
		t.root = nn
		return
	}

	current := t.root
	parent := &Node{
		iData:      0,
		fData:      0,
		leftChild:  nil,
		rightChild: nil,
	}

	for {
		parent = current
		if id < current.iData {
			current = current.leftChild
			if current == nil {
				parent.leftChild = nn
				return
			}
		} else {
			current = current.rightChild
			if current == nil {
				parent.rightChild = nn
				return
			}
		}
	}

}

func (t *Tree) delete(key int) bool {
	current := t.root
	parent := &Node{
		iData:      0,
		fData:      0,
		leftChild:  nil,
		rightChild: nil,
	}
	isLeftChild := true

	for current.iData != key {
		parent = current
		if key < current.iData {
			isLeftChild = true
			current = current.leftChild
		} else {
			isLeftChild = false
			current = current.rightChild
		}
		if current == nil {
			return false
		}
	}

	if current.leftChild == nil && current.rightChild == nil {
		if current == t.root {
			t.root = nil
		} else if isLeftChild {
			parent.leftChild = nil
		} else {
			parent.rightChild = nil
		}
	} else if current.rightChild == nil {
		if current == t.root {
			t.root = current.rightChild
		} else if isLeftChild {
			parent.leftChild = current.leftChild
		} else {
			parent.rightChild = current.leftChild
		}
	} else if current.leftChild == nil {
		if current == nil {
			t.root = current.rightChild
		} else if isLeftChild {
			parent.leftChild = current.rightChild
		} else {
			parent.rightChild = current.rightChild
		}
	} else {
		successor := getSuccessor(current)

		if current == t.root {
			t.root = successor
		} else if isLeftChild {
			parent.leftChild = successor
		} else {
			parent.rightChild = successor
			successor.leftChild = current.leftChild
		}
	}
	return true
}

func (t *Tree) inOrder(localRoot *Node) {
	if localRoot != nil {
		t.inOrder(localRoot.leftChild)

		localRoot.display()
		t.inOrder(localRoot.rightChild)
	}
}

func getSuccessor(dn *Node) *Node {
	successorParent := dn
	successor := dn
	current := dn.rightChild

	for current != nil {
		successorParent = successor
		successor = current
		current = current.leftChild
	}

	if successor != dn.rightChild {
		successorParent.leftChild = successor.rightChild
		successor.rightChild = dn.rightChild
	}

	return successor
}
