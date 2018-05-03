package tree

import (
	"fmt"
)

// Tree structure represents the classical binary tree
type Tree struct {
	Left  *Tree
	Right *Tree
	Value interface{}
}

// New creates tree with only root. The root value can be anything.
func New(root interface{}) *Tree {
	return &Tree{
		Left:  nil,
		Right: nil,
		Value: root,
	}
}

// SetChild add the tree a child. First sets the left, if nil, than sets the right.
// If both branches(left,right) are not nil, error returned.
func (t *Tree) SetChild(child *Tree) error {
	if t.Left == nil {
		t.Left = child
	} else if t.Right == nil {
		t.Right = child
	} else {
		return fmt.Errorf("tree has already assigned both children")
	}

	return nil
}
