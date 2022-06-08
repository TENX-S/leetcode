package daily

func minEatingSpeed(piles []int, h int) int {
	var max int
	for _, p := range piles {
		if p > max {
			max = p
		}
	}
	left, right := 1, max

	for left <= right {
		mid := left + (right-left)/2
		res := eatingHours(piles, mid)
		if res > h {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

func eatingHours(arr []int, k int) (res int) {
	for _, a := range arr {
		if a <= k {
			res += 1
		} else {
			if a%k == 0 {
				res += a / k
			} else {
				res += a/k + 1
			}
		}
	}
	return
}
