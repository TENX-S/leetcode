package Daily

import (
	"strconv"

	. "github.com/TENX-S/findwork/BinaryTree"
)

func findBottomLeftValue(root *TreeNode) int {
	var ans [][]string

	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		ans = append(ans, []string{})
		var p []*TreeNode
		for j := 0; j < len(q); j++ {
			node := q[j]
			if node == nil {
				ans[i] = append(ans[i], "null")
			} else {
				ans[i] = append(ans[i], strconv.Itoa(node.Val))
			}
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}

	for i, v := range ans[len(ans)-1] {
		if i%2 == 0 && v != "null" {
			res, _ := strconv.Atoi(v)
			return res
		}
	}

	return -1
}
