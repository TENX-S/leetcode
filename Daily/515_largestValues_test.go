package Daily

import (
	"github.com/TENX-S/findwork/BinaryTree"
	"testing"
)

func TestLargestValues(t *testing.T) {
	tree := BinaryTree.CreateBinaryTree("[0,-1]")
	t.Log(largestValues(tree))
}
