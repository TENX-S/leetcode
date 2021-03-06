package BinaryTree

func inorderTraversal(root *TreeNode) []int {
	var stack []*TreeNode
	var ans []int
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, node.Val)
		node = node.Right
	}

	return ans
}
