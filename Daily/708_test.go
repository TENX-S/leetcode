package Daily

import (
	"testing"
)

func TestInsert(t *testing.T) {
	node := CreateCircularLinkList([]int{1, 3, 5})
	insert(node, 4)
}
