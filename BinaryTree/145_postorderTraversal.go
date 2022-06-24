package BinaryTree

func postorderTraversal(root *TreeNode) []int {
	var stack []*TreeNode
	var ans []int
	var prev *TreeNode

	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Right == nil || root.Right == prev {
			ans = append(ans, root.Val)
			prev = root
			root = nil
		} else {
			stack = append(stack, root)
			root = root.Right
		}
	}

	return ans
}
