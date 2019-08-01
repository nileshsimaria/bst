// Package bst implements binary search tree data structure; purely for the educational purpose.
// It supports methods like Add one node, add bulk, walk. It does not support delete node.
package bst

import "fmt"

// Node type of the BST. It stores int data.
type Node struct {
	data  int
	left  *Node
	right *Node
}

// NewNode returns new BST node
func NewNode(data int, left *Node, right *Node) *Node {
	return &Node{
		data:  data,
		left:  left,
		right: right,
	}
}

// BST type
type BST struct {
	root *Node
}

// NewBST returns new empty BST
func NewBST() *BST {
	return &BST{
		root: nil,
	}
}

// AddBulk to add multiple data in the tree
func (t *BST) AddBulk(data ...int) error {
	for _, d := range data {
		if err := t.Add(d); err != nil {
			return err
		}
	}
	return nil
}

// Add into BST
func (t *BST) Add(data int) error {
	//time.Sleep(time.Second)
	if t.root == nil {
		t.root = NewNode(data, nil, nil)
		return nil
	}

	c := NewNode(data, nil, nil)
	p, err := findParent(t.root, data)
	if err != nil {
		return err
	}
	if p != nil {
		if p.data > data {
			p.left = c
		} else {
			p.right = c
		}
		return nil
	}

	return fmt.Errorf("Add(%d) failed: could not find parent", data)
}

// Search the node
func (t *BST) Search(data int) *Node {
	return search(t.root, data)
}

func search(r *Node, data int) *Node {
	if r.data == data {
		return r
	}
	if data < r.data && r.left != nil {
		return search(r.left, data)
	}
	if data > r.data && r.right != nil {
		return search(r.right, data)
	}
	return nil
}

func findParent(n *Node, data int) (*Node, error) {
	if n.data == data {
		return nil, fmt.Errorf("Found duplicate: data=%d", data)
	}
	if n.data > data && n.left == nil {
		return n, nil
	}

	if n.data < data && n.right == nil {
		return n, nil
	}

	if n.data > data && n.left != nil {
		return findParent(n.left, data)
	}

	if n.data < data && n.right != nil {
		return findParent(n.right, data)
	}

	return nil, nil
}

// Walk the tree in in-order fashion i.e. sort
// the data. Send back the data in the channel to
// the caller.
func (t *BST) Walk(dataC chan<- int) {
	walk(t.root, dataC)
	close(dataC)
}

func walk(n *Node, dataC chan<- int) {
	if n != nil {
		walk(n.left, dataC)
		dataC <- n.data
		walk(n.right, dataC)
	}
}
