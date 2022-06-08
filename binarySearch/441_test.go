package binarySearch

import "testing"

func TestArrangeCoins(t *testing.T) {
	for i := range []int{1, 1, 1, 1, 1, 1, 1} {
		i++
		t.Logf("%d -> %d", i, arrangeCoins(i))
	}
}
