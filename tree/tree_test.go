package tree

import (
  "testing"
)

func TestNew(t *testing.T) {
  var root interface{} = 5

  mytree := New(root)
  if mytree == nil {
    t.Error("Return from new should not be nil.")
  } else {
    if mytree.Value != root {
      t.Error("Root value is not correct assigned.")
    } else if mytree.Left != nil || mytree.Right != nil {
      t.Error("Left or Right branches is(are) not nil.")
    }
  }
}

func TestSetChild(t *testing.T) {
  mytree := New(5)
  leftChild := New(4)
  rightChild := New(6)

  if err := mytree.SetChild(leftChild); err != nil {
    t.Errorf("SetChild failed althought the tree children are empty.")
  } else if mytree.Left != leftChild {
    t.Errorf("SetChild hasn't assigned correct value.")
  }

  if err := mytree.SetChild(rightChild); err != nil {
    t.Errorf("SetChild failed althought the right children is empty.")
  } else if mytree.Right != rightChild {
    t.Errorf("SetChild hasn't assigned correct value.")
  }

  if err := mytree.SetChild(New(5)); err == nil {
    t.Errorf("SetChild should have failed?")
  }
}
