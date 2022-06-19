package BinaryTree

func preorderTraversal(root *TreeNode) []int {
	var stack []*TreeNode
	var ans []int
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			ans = append(ans, node.Val)
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return ans
}
