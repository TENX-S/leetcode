package Daily

import (
	"fmt"
	. "github.com/TENX-S/findwork/BinaryTree"
)

// 解法 1 两次DFS
func findFrequentTreeSum(root *TreeNode) []int {
	var m = make(map[int]int)
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		m[sumOfRoot(node)]++
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)

	fmt.Println(m)

	var max int
	for _, v := range m {
		if v > max {
			max = v
		}
	}
	fmt.Println(max)

	var ans []int
	for k, v := range m {
		if v == max {
			ans = append(ans, k)
		}
	}

	return ans
}

func sumOfRoot(root *TreeNode) int {
	var res int
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		res += node.Val
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return res
}

// 解法 2 一次后序遍历

func findFrequentTreeSum2(root *TreeNode) []int {
	m := make(map[int]int)
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftSum := dfs(node.Left)
		rightSum := dfs(node.Right)
		rootSum := leftSum + rightSum + node.Val
		m[rootSum]++
		return rootSum
	}
	dfs(root)

	var max int
	for _, v := range m {
		if v > max {
			max = v
		}
	}

	var ans []int
	for k, v := range m {
		if v == max {
			ans = append(ans, k)
		}
	}

	return ans
}
