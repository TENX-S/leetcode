package BinarySearch

func judgeSquareSum(c int) bool {
	for b := mySqrt(c); b >= 0; b-- {
		if isPerfectSquare(c - b*b) {
			return true
		}
	}
	return false
}
