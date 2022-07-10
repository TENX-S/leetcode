package SlidingWindow

import "strconv"

func divisorSubstrings(num, k int) (ans int) {
	n := strconv.Itoa(num)
	left, right := 0, 0

	for right <= len(n) {
		if right-left < k {
			right++
			continue
		}
		val, _ := strconv.Atoi(n[left:right])
		if val == 0 {
			left++
			continue
		}
		if num%val == 0 {
			ans++
		}
		left++
	}
	return
}
