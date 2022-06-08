package binarytree

var res int

func rangeSumBST(root *TreeNode, low, high int) int {
	traverse(root, low, high)
	return res
}

func traverse(root *TreeNode,low,high int) {
	if root == nil {
		return
	}
	curr := root.Val
	if curr >= low && curr <= high {
		res += curr
	}
	traverse(root.Left,low,high)
	traverse(root.Right,low,high)
}

