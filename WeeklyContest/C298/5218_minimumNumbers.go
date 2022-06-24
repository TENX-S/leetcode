package C298

import "strconv"

func minimumNumbers(num int, k int) int {
	if num == 0 {
		return 0
	}

	if k == getLastNum(num) {
		return 1
	}

	if k == 0 && num%10 != 0 {
		return -1
	}

	if k%2 == 0 && num%2 != 0 {
		return -1
	}

	var ans int
	for num > k {
		ans++
		num -= k
		if num > 9 && getLastNum(num) == k {
			return ans + 1
		}
	}
	if num == k {
		return ans + 1
	}

	return -1
}

func getLastNum(num int) int {
	s := strconv.Itoa(num)
	last, _ := strconv.Atoi(string(s[len(s)-1]))
	return last
}
