package BinaryTree

import (
	"testing"
)

func Test_binaryTreePaths2(t *testing.T) {
	tree := CreateBinaryTree("[1,2,3,null,5,null,null]")
	binaryTreePaths2(tree)
}
