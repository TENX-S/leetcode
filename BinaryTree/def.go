package BinaryTree

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const (
	Preorder = iota
	Inorder
	Postorder
)

func createBinaryTree(arr string) *TreeNode {
	arr = strings.TrimPrefix(arr, "[")
	arr = strings.TrimSuffix(arr, "]")
	strTree := strings.Split(arr, ",")
	arrTree := make([]*TreeNode, len(strTree))

	var root *TreeNode
	for i, data := range strTree {
		var node *TreeNode
		if data != "null" {
			val, err := strconv.Atoi(data)
			if err != nil {
				fmt.Println("Invalid value: ", data)
			}
			node = &TreeNode{Val: val}
		}
		arrTree[i] = node
		if i == 0 {
			root = node
		}
	}

	for i := 0; 2*i+2 < len(strTree); i++ {
		if arrTree[i] != nil {
			arrTree[i].Left = arrTree[2*i+1]
			arrTree[i].Right = arrTree[2*i+2]
		}
	}

	return root
}

func (node *TreeNode) PrintBinaryTree(order int) {
	var traversalArr []int
	var traverse func(*TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		if order == Preorder {
			traversalArr = append(traversalArr, node.Val)
		}
		traverse(node.Left)
		if order == Inorder {
			traversalArr = append(traversalArr, node.Val)
		}
		traverse(node.Right)
		if order == Postorder {
			traversalArr = append(traversalArr, node.Val)
		}
	}
	traverse(node)

	fmt.Println(traversalArr)
}
