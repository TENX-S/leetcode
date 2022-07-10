package C301

import "testing"

func TestSmallestInfiniteSet(t *testing.T) {
	s := Constructor()
	t.Log(s.PopSmallest())
	s.AddBack(1)
	t.Log(s.PopSmallest())
	t.Log(s.PopSmallest())
	t.Log(s.PopSmallest())
	s.AddBack(2)
	s.AddBack(3)
	t.Log(s.PopSmallest())
	t.Log(s.PopSmallest())
}
