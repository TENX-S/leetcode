package C301

func canChange(start, target string) bool {
	j := 0
	for i := 0; i < len(start); i++ {
		if start[i] == byte('_') {
			continue
		}
		for j < len(target) && target[j] == byte('_') {
			j++
		}
		if j == len(start) {
			return false
		}
		if start[i] != target[j] {
			return false
		}
		if start[i] == byte('L') && i < j {
			return false
		}
		if start[i] == byte('R') && i > j {
			return false
		}
		j++
	}
	for i := j; i < len(start); i++ {
		if target[i] != byte('_') {
			return false
		}
	}

	return true
}
