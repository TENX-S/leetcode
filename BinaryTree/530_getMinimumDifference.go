package BinaryTree

import "math"

func getMinimumDifference(root *TreeNode) (ans int) {
	ans = math.MaxInt64
	if root == nil {
		return
	}
	prev := -1
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if prev != -1 && node.Val-prev < ans {
			ans = node.Val - prev
		}
		prev = node.Val
		dfs(node.Right)
	}
	dfs(root)

	return
}
