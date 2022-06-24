package BinaryTree

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	var dfs func(*TreeNode)
	var ans bool
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if isSameTree(node, subRoot) {
			ans = true
		}
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)

	return ans
}
