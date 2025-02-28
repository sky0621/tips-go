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

func (n *TreeNode) String(indent string) string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("{\n"))
	sb.WriteString(fmt.Sprintf("%s  Value: %d, \n", indent, n.Value))
	if n.Left != nil {
		sb.WriteString(fmt.Sprintf("%s  Left: %+v, \n", indent, n.Left.String(indent+"  ")))
	}
	if n.Right != nil {
		sb.WriteString(fmt.Sprintf("%s  Right: %+v, \n", indent, n.Right.String(indent+"  ")))
	}
	sb.WriteString(fmt.Sprintf("%s}", indent))
	return sb.String()
}

func main() {
	root := &TreeNode{
		Value: 1,
		Left: &TreeNode{
			Value: 11,
			Left: &TreeNode{
				Value: 111,
				Left: &TreeNode{
					Value: 1111,
					Left: &TreeNode{
						Value: 11111,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
				Right: &TreeNode{
					Value: 1112,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &TreeNode{
				Value: 112,
				Left:  nil,
				Right: &TreeNode{
					Value: 1122,
					Left:  nil,
					Right: nil,
				}},
		},
		Right: &TreeNode{
			Value: 12,
			Left: &TreeNode{
				Value: 121,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Value: 122,
				Left:  nil,
				Right: nil,
			},
		},
	}
	fmt.Println(root.String(""))
}
