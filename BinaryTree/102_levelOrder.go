package BinaryTree

func levelOrder(root *TreeNode) [][]int {
	var ans [][]int
	if root == nil {
		return ans
	}

	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		ans = append(ans, []int{})
		var p []*TreeNode
		for j := 0; j < len(q); j++ {
			node := q[j]
			ans[i] = append(ans[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}

	return ans
}
