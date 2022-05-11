package linklist

import "testing"

func TestSplitListToParts(t *testing.T) {
	res := splitListToParts(NewFromList([]int{}), 3)
	for _, r := range res {
		r.PrintList()
	}
}
