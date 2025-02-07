package main

import (
	"fmt"
	"strings"
)

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func (n *TreeNode) String() string {
	if n == nil {
		return "<nil>"
	}
	sb := strings.Builder{}
	sb.WriteString("TreeNode {\n")
	sb.WriteString(fmt.Sprintf("Value: %d, \n", n.Value))
	sb.WriteString(fmt.Sprintf("Left: %+v, \n", n.Left.String()))
	sb.WriteString(fmt.Sprintf("Right: %+v, \n", n.Right.String()))
	sb.WriteString("}")
	return sb.String()
}

func main() {
	root := &TreeNode{
		Value: 10,
		Left: &TreeNode{
			Value: 5,
			Left:  &TreeNode{Value: 3},
			Right: &TreeNode{Value: 20},
		},
		Right: &TreeNode{
			Value: 15,
			Left:  &TreeNode{Value: 7},
			Right: &TreeNode{Value: 11},
		},
	}
	fmt.Println(root.String())
}
