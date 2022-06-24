package SfTech

import "testing"

func TestHashCycle(t *testing.T) {
	t.Log(hasCycle("1->4,2->5,3->6,3->7,4->8,5->8,5->9,6->9,6->11,7->11,8->12,9->12,9->13,10->6,10->13,10->14,11->10,11->14"))
}
