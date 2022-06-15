package BinaryTree

func diameterOfBinaryTree(root *TreeNode) int {
	maxDia := 0

	var dfs func(*TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftMax := dfs(root.Left)
		rightMax := dfs(root.Right)
		maxDia = Max(leftMax+rightMax, maxDia)

		return 1 + Max(leftMax, rightMax)
	}

	dfs(root)

	return maxDia
}

func Max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
