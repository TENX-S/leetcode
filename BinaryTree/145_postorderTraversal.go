package BinaryTree

func postorderTraversal(root *TreeNode) []int {
	var stack []*TreeNode
	var ans []int

	node := root

	for node != nil || len(stack) > 0 {
		stack = append(stack, node)
		node = node.Left
	}

	node = stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	node = node.Right
	ans = append(ans, node.Val)
	return ans
}
