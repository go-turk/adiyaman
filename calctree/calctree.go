package calctree

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Node struct {
	Value int   `json:"value"`
	Left  *Node `json:"left"`
	Right *Node `json:"right"`
}

func NewNode(value int) *Node {
	return &Node{Value: value}
}

func (n *Node) AddChildToLeft(child *Node) {
	n.Left = child
}

func (n *Node) AddChildToRight(child *Node) {
	n.Right = child
}

func (n *Node) SharewithChildren() {
	if n.Value <= 10 {
		return
	}
	rnd := rand.New(rand.NewSource(time.Now().Unix()))
	leftValue := rnd.Intn(n.Value-2) + 1
	rightValue := n.Value - leftValue
	n.AddChildToLeft(NewNode(leftValue))
	n.Left.SharewithChildren()
	n.AddChildToRight(NewNode(rightValue))
	n.Right.SharewithChildren()
}

func (n *Node) ForceToTree(rate int) int {
	total := 0
	if n.Right == nil || n.Left == nil {
		return 0
	}
	total += n.Right.ForceToTree(rate)
	total += n.Left.ForceToTree(rate)
	rnd := rand.New(rand.NewSource(time.Now().Unix()))
	if rnd.Intn(100) > (100 - rate) {
		if number := rnd.Intn(n.Value + 1); number != n.Value {
			total++
			n.Value = number
		}
	}
	return total
}

func (n *Node) CheckHealth() {
	if n.Right == nil || n.Left == nil {
		return
	}
	if n.Right.Value+n.Left.Value != n.Value {
		fmt.Println("Error: ", n.Value)
	}
	n.Right.CheckHealth()
	n.Left.CheckHealth()
}

func (n *Node) Repair() int {
	total := 0
	if n.Right == nil || n.Left == nil {
		return 0
	}
	total += n.Right.Repair()
	total += n.Left.Repair()
	if val := n.Right.Value + n.Left.Value; val != n.Value {
		total++
		n.Value = val
	}
	return total
}

func (n *Node) BinaryTreeToJSON() string {
	if n.Right == nil || n.Left == nil {
		return fmt.Sprintf("{\"value\":%d}", n.Value)
	}
	return fmt.Sprintf("{\"value\":%d, \"left\":%s, \"right\":%s}", n.Value, n.Left.BinaryTreeToJSON(), n.Right.BinaryTreeToJSON())
}

func (n *Node) SaveTree(filename string) error {
	result := n.BinaryTreeToJSON()
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	file.WriteString(result)
	return nil
}

func LoadTree(filename string) (*Node, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var root Node
	err = json.Unmarshal(file, &root)
	return &root, nil
}
