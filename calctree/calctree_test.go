package calctree_test

import (
	"github.com/go-turk/adiyaman/calctree"
	"testing"
)

func TestNewNode(t *testing.T) {
	newNode := calctree.NewNode(50)
	if newNode.Value != 50 {
		t.Error("NewNode value is not 10")
	}
	if newNode.Left != nil {
		t.Error("NewNode left is not nil")
	}
	if newNode.Right != nil {
		t.Error("NewNode right is not nil")
	}
}

func TestAddChildToLeft(t *testing.T) {
	newNode := calctree.NewNode(50)
	newNode.AddChildToLeft(calctree.NewNode(20))
	if newNode.Left.Value != 20 {
		t.Error("NewNode left is not 20")
	}
}

func TestAddChildToRight(t *testing.T) {
	newNode := calctree.NewNode(50)
	newNode.AddChildToRight(calctree.NewNode(20))
	if newNode.Right.Value != 20 {
		t.Error("NewNode right is not 20")
	}
}

func TestSharewithChildren(t *testing.T) {
	newNode := calctree.NewNode(50)
	newNode.SharewithChildren()
	if newNode.Left.Value+newNode.Right.Value != 50 {
		t.Error("NewNode left and right values are not 50")
	}
}

func TestForceToTree(t *testing.T) {
	newNode := calctree.NewNode(50)
	newNode.SharewithChildren()
	newNode.ForceToTree(100)
	if newNode.Left.Value+newNode.Right.Value == 50 {
		t.Error("NewNode left and right values are 50")
	}
}

func TestNode_Repair(t *testing.T) {
	newNode := calctree.NewNode(100)
	newNode.SharewithChildren()
	newNode.ForceToTree(100)
	newNode.Repair()
	if newNode.Left.Value+newNode.Right.Value != 100 {
		t.Errorf("NewNode left and right values are not 20 is %d", newNode.Left.Value+newNode.Right.Value)
	}
}

func FuzzNode_Repair(f *testing.F) {
	f.Add(100)
	f.Fuzz(func(t *testing.T, value int) {
		newNode := calctree.NewNode(value)
		newNode.SharewithChildren()
		newNode.ForceToTree(100)
		newNode.Repair()
		if newNode.Left == nil || newNode.Right == nil {
			return
		}
		if newNode.Left.Value+newNode.Right.Value != value {
			f.Errorf("NewNode left and right values are not 20 is %d", newNode.Left.Value+newNode.Right.Value)
		}
	})
}
