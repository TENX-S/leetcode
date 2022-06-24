package BinaryTree

func sumOfLeftLeaves(root *TreeNode) int {
	var dfs func(*TreeNode)
	var ans int
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
			ans += root.Left.Val
		}
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return ans
}
