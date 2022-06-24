package BinaryTree

import (
	"strconv"
	"strings"
)

func binaryTreePaths(root *TreeNode) []string {
	var dfs func(*TreeNode)
	var ans []string
	var path []string
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		if node.Left == nil && node.Right == nil {
			path = append(path, strconv.Itoa(node.Val))
			ans = append(ans, strings.Join(path, "->"))
			path = path[:len(path)-1]
		}
		path = append(path, strconv.Itoa(node.Val))
		dfs(node.Left)
		dfs(node.Right)
		path = path[:len(path)-1]
	}
	dfs(root)

	return ans
}

func binaryTreePaths2(root *TreeNode) []string {
	var dfs func(*TreeNode, string)
	var res []string
	var path string
	dfs = func(node *TreeNode, path string) {
		path += strconv.Itoa(node.Val)
		if node.Left == nil && node.Right == nil {
			res = append(res, path)
			return
		}
		if node.Left != nil {
			dfs(node.Left, path+"->")
		}
		if node.Right != nil {
			dfs(node.Right, path+"->")
		}
	}
	dfs(root, path)
	return res
}
