package C298

func greatestLetter(s string) string {
	var L []byte
	for _, l := range []byte(s) {
		if l <= 90 {
			L = append(L, l)
		}
	}
	var ans []byte
	for _, n := range L {
		for _, r := range []byte(s) {
			if r >= 97 && r-n == 32 {
				ans = append(ans, n)
			}
		}
	}

	if len(ans) == 0 {
		return ""
	}
	max := ans[0]
	for _, v := range ans {
		if v > max {
			max = v
		}
	}
	return string(max)
}
