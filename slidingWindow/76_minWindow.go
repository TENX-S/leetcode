package slidingWindow

import "math"

func minWindow(s, t string) string {
	need := make(map[byte]int, len(t))
	window := make(map[byte]int, len(t))

	for _, c := range []byte(t) {
		need[c]++
	}

	left, right := 0, 0
	valid := 0
	start, length := 0, math.MaxInt32
	for right < len(s) {
		c := s[right]
		right++
		if cnt, exist := need[c]; exist {
			window[c]++
			if window[c] == cnt {
				valid++
			}
		}

		for valid == len(need) {
			if right-left < length {
				start = left
				length = right - left
			}
			d := s[left]
			left++
			if cnt, exist := need[d]; exist {
				if window[d] == cnt {
					valid--
				}
				window[d]--
			}
		}
	}
	if length == math.MaxInt32 {
		return ""
	} else {
		return s[start : start+length]
	}
}
