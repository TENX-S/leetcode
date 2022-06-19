package Daily

import (
	"testing"
)

func TestFindFrequentTreeSum(t *testing.T) {
	tree := CreateBinaryTree([]int{5, 2, -3})
	findFrequentTreeSum(tree)
}
