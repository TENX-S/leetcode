package BinaryTree

import (
	"testing"
)

func TestCreateBinaryTree(t *testing.T) {
	tree := createBinaryTree("[4,1,6,0,2,5,7,null,null,null,3,null,null,null,8]")
	tree.PrintBinaryTree(Preorder)
	tree.PrintBinaryTree(Inorder)
	tree.PrintBinaryTree(Postorder)
}
