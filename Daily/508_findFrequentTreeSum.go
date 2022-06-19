package Daily

import "fmt"

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
	for k := range m {
		if k > max {
			max = k
		}
	}

	var ans []int
	for k, v := range m {
		if k == max {
			ans = append(ans, v)
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
