package BinarySearch

func arrangeCoins(n int) int {
	budget := func(x int) int {
		return x * (x + 1) / 2
	}

	left, right := 1, n/2
	for left <= right {
		mid := left + (right-left)/2
		key := budget(mid)
		if key == n {
			left = mid + 1
		} else if key < n {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if budget(left) > n {
		return left - 1
	}

	return left
}
