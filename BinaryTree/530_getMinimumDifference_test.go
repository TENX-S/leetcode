package BinaryTree

import "testing"

func TestGetMinimumDifference(t *testing.T) {
	tree := CreateBinaryTree("[4,2,6,1,3,null,null]")
	getMinimumDifference(tree)
}
