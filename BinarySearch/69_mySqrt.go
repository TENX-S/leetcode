package BinarySearch

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	if x <= 2 {
		return 1
	}

	left, right := 1, x/2
	for left <= right {
		mid := left + (right-left)/2
		key := mid * mid
		if key == x {
			return mid
		} else if key < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if left*left > x {
		return left - 1
	}
	return left
}
