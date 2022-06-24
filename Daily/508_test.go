package Daily

import (
	. "github.com/TENX-S/findwork/BinaryTree"
	"testing"
)

func TestFindFrequentTreeSum(t *testing.T) {
	tree := CreateBinaryTree("[5,2,-3]")
	t.Log(findFrequentTreeSum2(tree))
}
