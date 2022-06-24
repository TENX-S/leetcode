package Daily

import (
	. "github.com/TENX-S/findwork/BinaryTree"
	"math"
)

func largestValues(root *TreeNode) (ans []int) {
	res := levelOrder(root)
	for _, arr := range res {
		ans = append(ans, max(arr))
	}
	return
}

func max(arr []int) int {
	cand := math.MaxInt64
	for _, a := range arr {
		if a > cand {
			cand = a
		}
	}
	return cand
}

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
