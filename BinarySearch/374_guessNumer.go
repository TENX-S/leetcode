package BinarySearch

func guess(num int) int {
	return -1
}

func guessNumber(n int) int {
	left, right := 1, n
	mid := (right + left) / 2

	for {
		switch guess(mid) {
		case -1:
			right = mid - 1
			mid = (right + left) / 2
		case 0:
			return mid
		case 1:
			left = mid + 1
			mid = (right + left) / 2
		}
	}
}
