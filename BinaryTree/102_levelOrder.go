package BinaryTree

func levelOrder(root *TreeNode) (res [][]int) {
	if root == nil {
		return
	}

	//本层节点
	curr := []*TreeNode{root}
	//本层节点为空时退出
	for i := 0; len(curr) > 0; i++ {
		//i+1代表层级，0为第一层
		//不为空，为本层节点新建一个数组存放节点值
		res = append(res, []int{})
		var next []*TreeNode
		//在加入本层节点值的同时进行拓展得到下层节点
		for j := 0; j < len(curr); j++ {
			node := curr[j]
			res[i] = append(res[i], node.Val)
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		curr = next
	}

	return
}
