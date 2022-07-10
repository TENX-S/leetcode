package C300

import "testing"
import . "github.com/TENX-S/findwork/LinkedList"

func Test_spiralMatrix(t *testing.T) {
	head := NewFromList([]int{3, 0, 2, 6, 8, 1, 7, 9, 4, 2, 5, 5, 0})
	t.Log(head.Val)
	t.Log(spiralMatrix(3, 5, head))
}
