package binarytree

import (
	"fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

const (
	PreOrder = iota
	InfixOrder
	PostOrder
)

// func NewTree(arr []int, order int) *TreeNode {
// 
// 	node := &TreeNode{Val: }
// }

func PrintTree(root *TreeNode, order int) {
	if root == nil {
		return;
	}
	if order == PreOrder {
		fmt.Println(root.Val)
	}
	PrintTree(root.Left, order)
	if order == InfixOrder {
		fmt.Println(root.Val)
	}
	PrintTree(root.Right, order)
	if order == PostOrder {
		fmt.Println(root.Val)
	}
}

